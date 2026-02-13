package tui

import (
	"strings"
	"time"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/model"
	"shop.gyeongho.dev/internal/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// loadCompleteMsg is sent when the loading phase ends.
type loadCompleteMsg struct{}

// Page identifies the current TUI section.
type Page int

const (
	PageLanding Page = iota
	PageShop
	PageAbout
	PageFAQ
	PageCart
)

func (p Page) String() string {
	switch p {
	case PageLanding:
		return "Landing"
	case PageShop:
		return "Shop"
	case PageAbout:
		return "About"
	case PageFAQ:
		return "FAQ"
	case PageCart:
		return "Cart"
	default:
		return "Unknown"
	}
}

// Model is the root Bubble Tea model for the TUI.
type Model struct {
	Client       apiclient.Client
	CurrentPage  Page
	Width        int
	Height       int
	ScrollOffset int
	Cursor       int
	Products     []model.Product
	Cart         model.Cart
	AddQuantity  int // quantity to add on next Enter (Shop page); min 1
	About        model.StoreInfo
	FAQ          []model.FAQEntry
	Loading      bool
	Error        string
}

// NewModel returns an initial model with the given API client.
func NewModel(client apiclient.Client) Model {
	return Model{
		Client:      client,
		CurrentPage: PageLanding,
		Width:       80,
		Height:      24,
		Loading:     true,
		Cart:        model.Cart{},
		AddQuantity: 1,
	}
}

// Init runs once at start and returns a command that sends loadCompleteMsg after a delay.
func (m Model) Init() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return loadCompleteMsg{}
	})
}

// Update handles key events, window resize, and load-complete messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case KeyCtrlC, KeyQuit:
			return m, tea.Quit
		case KeyShop:
			m.CurrentPage = PageShop
			m.ScrollOffset, m.Cursor = 0, 0
			m.AddQuantity = 1
			return m, nil
		case KeyAbout:
			m.CurrentPage = PageAbout
			m.ScrollOffset = 0
			return m, nil
		case KeyFAQ:
			m.CurrentPage = PageFAQ
			m.ScrollOffset = 0
			return m, nil
		case KeyCart:
			m.CurrentPage = PageCart
			m.ScrollOffset = 0
			m.Cursor = 0
			return m, nil
		case KeyUp:
			if m.CurrentPage == PageShop && m.Cursor > 0 {
				m.Cursor--
				m.AddQuantity = 1
			} else if m.CurrentPage == PageCart && len(m.Cart.Items) > 0 && m.Cursor > 0 {
				m.Cursor--
			} else if m.ScrollOffset > 0 {
				m.ScrollOffset--
			}
			return m, nil
		case KeyDown:
			if m.CurrentPage == PageShop && m.Cursor < len(m.Products)-1 {
				m.Cursor++
				m.AddQuantity = 1
			} else if m.CurrentPage == PageCart && len(m.Cart.Items) > 0 && m.Cursor < len(m.Cart.Items)-1 {
				m.Cursor++
			} else {
				m.ScrollOffset++
			}
			return m, nil
		case KeyPlus:
			if m.CurrentPage == PageShop {
				m.AddQuantity++
			}
			return m, nil
		case KeyMinus:
			if m.CurrentPage == PageCart && len(m.Cart.Items) > 0 {
				m.Cart.DecreaseQuantity(m.Cursor)
				if len(m.Cart.Items) > 0 && m.Cursor >= len(m.Cart.Items) {
					m.Cursor = len(m.Cart.Items) - 1
				} else {
					m.Cursor = 0
				}
			} else if m.CurrentPage == PageShop && m.AddQuantity > 1 {
				m.AddQuantity--
			}
			return m, nil
		case KeyBackspace:
			if m.CurrentPage == PageCart && len(m.Cart.Items) > 0 {
				m.Cart.RemoveItem(m.Cursor)
				if len(m.Cart.Items) > 0 && m.Cursor >= len(m.Cart.Items) {
					m.Cursor = len(m.Cart.Items) - 1
				} else {
					m.Cursor = 0
				}
			}
			return m, nil
		case KeyEnter:
			if m.CurrentPage == PageShop && len(m.Products) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Products) {
				p := m.Products[m.Cursor]
				if p.Quantity > 0 {
					qty := m.AddQuantity
					if qty > p.Quantity {
						qty = p.Quantity
					}
					m.Cart.AddOrMergeItem(p.ID, p.Name, p.Price, qty, p.Quantity)
					m.AddQuantity = 1
				}
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case loadCompleteMsg:
		m.Loading = false
		return m, nil
	}
	return m, nil
}

const (
	maxContentWidth = 60
	headerLines     = 3
	footerLines     = 2
	maxBodyHeight   = 40
)

// View returns the current frame: a centered loading view or a centered main view
// (header, body, footer). Layout is responsive to Width and Height; output is
// clipped to the viewport height so the menu is never cut off.
func (m Model) View() string {
	w, h := m.Width, m.Height
	if w <= 0 {
		w = 80
	}
	if h <= 0 {
		h = 24
	}
	if m.Loading {
		loadingView := Loader()
		return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, loadingView)
	}
	header := RenderHeader(m.CurrentPage, m.Cart.Total(), m.Cart.ItemCount())
	footer := RenderFooter()
	bodyWidth := w
	if bodyWidth > maxContentWidth {
		bodyWidth = maxContentWidth
	}
	var body string
	switch m.CurrentPage {
	case PageLanding:
		body = pages.Landing()
	case PageShop:
		body = pages.Shop(m.Products, m.ScrollOffset, m.Cursor, bodyWidth, m.AddQuantity)
	case PageAbout:
		body = pages.About(m.About)
	case PageFAQ:
		body = pages.FAQ(m.FAQ, bodyWidth)
	case PageCart:
		body = pages.Cart(m.Cart, m.Cursor)
	default:
		body = pages.Landing()
	}
	if m.Error != "" {
		body = "Error: " + m.Error + "\n\nPress a/s/d to continue."
	}
	body = lipgloss.NewStyle().Width(bodyWidth).MaxWidth(bodyWidth).Render(body)
	available := h - headerLines - footerLines
	mainVerticalPad := 0
	if available > 14 {
		mainVerticalPad = 2
	}
	if available > 24 {
		mainVerticalPad = 4
	}
	bodyHeight := available - 2*mainVerticalPad
	if bodyHeight < 1 {
		bodyHeight = 1
	}
	if bodyHeight > maxBodyHeight {
		bodyHeight = maxBodyHeight
	}
	body = bodyViewport(body, m.ScrollOffset, bodyHeight)
	mainView := header + "\n" + body + "\n" + footer
	if mainVerticalPad > 0 {
		mainView = strings.Repeat("\n", mainVerticalPad) + mainView + strings.Repeat("\n", mainVerticalPad)
	}
	mainViewLines := strings.Split(mainView, "\n")
	if len(mainViewLines) > h {
		mainViewLines = mainViewLines[:h]
		mainView = strings.Join(mainViewLines, "\n")
	}
	return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, mainView)
}

// bodyViewport returns exactly height lines: a window into content at scrollOffset, padded with blank lines if needed.
func bodyViewport(content string, scrollOffset, height int) string {
	lines := strings.Split(content, "\n")
	if height <= 0 {
		return content
	}
	maxScroll := len(lines) - height
	if maxScroll < 0 {
		maxScroll = 0
	}
	start := scrollOffset
	if start > maxScroll {
		start = maxScroll
	}
	if start < 0 {
		start = 0
	}
	end := start + height
	if end > len(lines) {
		end = len(lines)
	}
	visible := lines[start:end]
	for len(visible) < height {
		visible = append(visible, "")
	}
	return strings.Join(visible, "\n")
}
