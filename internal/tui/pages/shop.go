package pages

import (
	"fmt"
	"strings"

	"mushroom.gyeongho.dev/internal/model"

	"github.com/charmbracelet/lipgloss"
)

var (
	shopNameStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	shopAttrStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	shopPriceStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5A00"))
	shopDescStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	shopQuantityStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
)

// Shop renders the shop page: product name, attributes, price, description, and quantity per item.
func Shop(products []model.Product, scrollOffset, cursor, width int) string {
	if len(products) == 0 {
		return "No products. Press a/s/d to navigate."
	}
	var b strings.Builder
	for i, p := range products {
		cursorMark := " "
		if i == cursor {
			cursorMark = ">"
		}
		b.WriteString(shopNameStyle.Render(cursorMark + " " + p.Name))
		b.WriteString("\n")
		if len(p.Attributes) > 0 {
			b.WriteString(shopAttrStyle.Render(strings.Join(p.Attributes, " | ")))
			b.WriteString("\n")
		}
		b.WriteString(shopPriceStyle.Render(fmt.Sprintf("$%.2f", p.Price)))
		b.WriteString("\n")
		b.WriteString(shopDescStyle.Render(p.Description))
		b.WriteString("\n")
		qty := "Sold out!"
		if p.Quantity > 0 {
			qty = fmt.Sprintf("%d left", p.Quantity)
		}
		b.WriteString(shopQuantityStyle.Render(qty))
		b.WriteString("\n\n")
	}
	b.WriteString("Enter=add to cart  c=cart  a/s/d=nav")
	return b.String()
}
