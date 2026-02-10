package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const navWidth = 10

func centerText(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	pad := (width - len(s)) / 2
	return strings.Repeat(" ", pad) + s + strings.Repeat(" ", width-pad-len(s))
}

// RenderHeader returns the header (terminal.shop style: logo, a shop, s about, d faq, cart).
func RenderHeader(currentPage Page, cartTotal float64, cartCount int) string {
	cartStr := centerText(fmt.Sprintf("cart $%.2f [%d]", cartTotal, cartCount), 20)
	parts := []struct {
		label string
		width int
	}{
		{centerText("mushroom", navWidth), navWidth},
		{centerText("a shop", navWidth), navWidth},
		{centerText("s about", navWidth), navWidth},
		{centerText("d faq", navWidth), navWidth},
		{cartStr, 20},
	}
	bold := lipgloss.NewStyle().Bold(true)
	top := "┌"
	bot := "└"
	line := ""
	for i, p := range parts {
		top += strings.Repeat("─", p.width)
		bot += strings.Repeat("─", p.width)
		if i < len(parts)-1 {
			top += "┬"
			bot += "┴"
		}
		cell := p.label
		if (i == 1 && currentPage == PageShop) || (i == 2 && currentPage == PageAbout) || (i == 3 && currentPage == PageFAQ) {
			cell = bold.Render(cell)
		}
		line += cell
		if i < len(parts)-1 {
			line += "│"
		}
	}
	top += "┐\n"
	bot += "┘"
	return top + "│" + line + "│\n" + bot
}
