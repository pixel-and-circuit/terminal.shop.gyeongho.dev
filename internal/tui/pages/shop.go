package pages

import (
	"fmt"
	"mushroom.gyeongho.dev/internal/model"
)

// Shop renders the shop page: product list with scroll and cursor.
func Shop(products []model.Product, scrollOffset, cursor, width int) string {
	if len(products) == 0 {
		return "No products. Press a/s/d to navigate."
	}
	var s string
	for i, p := range products {
		cursorMark := " "
		if i == cursor {
			cursorMark = ">"
		}
		qty := "Sold out!"
		if p.Quantity > 0 {
			qty = fmt.Sprintf("%d left", p.Quantity)
		}
		s += fmt.Sprintf("%s %s - $%.2f (%s)\n  %s\n", cursorMark, p.Name, p.Price, qty, p.Description)
	}
	return s + "\nEnter=add to cart  c=cart  a/s/d=nav"
}
