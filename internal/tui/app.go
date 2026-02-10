package tui

import (
	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/model"
	"mushroom.gyeongho.dev/internal/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
)

// Page is the current TUI section.
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

// Model is the root Bubble Tea model.
type Model struct {
	Client       apiclient.Client
	CurrentPage  Page
	Width        int
	Height       int
	ScrollOffset int
	Cursor       int
	Products     []model.Product
	Cart         model.Cart
	About        model.StoreInfo
	FAQ          []model.FAQEntry
	Loading      bool
	Error        string
}

// NewModel returns an initial model (inject client from main).
func NewModel(client apiclient.Client) Model {
	return Model{
		Client:      client,
		CurrentPage: PageLanding,
		Width:       60,
		Height:      24,
		Cart:        model.Cart{},
	}
}

// Init runs once at start; can trigger loading.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages (key, window size).
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case KeyCtrlC, KeyQuit:
			return m, tea.Quit
		case KeyShop:
			m.CurrentPage = PageShop
			m.ScrollOffset, m.Cursor = 0, 0
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
			return m, nil
		case KeyUp:
			if m.CurrentPage == PageShop && m.Cursor > 0 {
				m.Cursor--
			} else if m.CurrentPage == PageFAQ && m.ScrollOffset > 0 {
				m.ScrollOffset--
			}
			return m, nil
		case KeyDown:
			if m.CurrentPage == PageShop && m.Cursor < len(m.Products)-1 {
				m.Cursor++
			} else if m.CurrentPage == PageFAQ && m.ScrollOffset < len(m.FAQ)-1 {
				m.ScrollOffset++
			}
			return m, nil
		case KeyEnter:
			if m.CurrentPage == PageShop && len(m.Products) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Products) {
				p := m.Products[m.Cursor]
				if p.Quantity > 0 {
					m.Cart.Items = append(m.Cart.Items, model.CartItem{
						ProductID: p.ID, Name: p.Name, UnitPrice: p.Price, Quantity: 1,
					})
				}
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	}
	return m, nil
}

// View renders the TUI: header + body (by page) + footer.
func (m Model) View() string {
	header := RenderHeader(m.CurrentPage, m.Cart.Total(), m.Cart.ItemCount())
	footer := RenderFooter()
	var body string
	switch m.CurrentPage {
	case PageLanding:
		body = pages.Landing()
	case PageShop:
		body = pages.Shop(m.Products, m.ScrollOffset, m.Cursor, m.Width)
	case PageAbout:
		body = pages.About(m.About)
	case PageFAQ:
		body = pages.FAQ(m.FAQ, m.ScrollOffset, m.Width)
	case PageCart:
		body = pages.Cart(m.Cart)
	default:
		body = pages.Landing()
	}
	if m.Error != "" {
		body = "Error: " + m.Error + "\n\nPress a/s/d to continue."
	}
	return header + "\n" + body + "\n" + footer
}
