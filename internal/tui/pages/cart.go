package pages

import (
	"fmt"
	"shop.gyeongho.dev/internal/model"
)

// Cart renders the cart page: items (with cursor on selectedIndex), total, action buttons, and hints.
func Cart(c model.Cart, selectedIndex int) string {
	if len(c.Items) == 0 {
		return "Cart is empty.\n\n[ Keep shopping ]  (a)\n\nPress a to go to shop."
	}
	var s string
	for j, i := range c.Items {
		cursor := "  "
		if j == selectedIndex {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s x%d ₩%d\n", cursor, i.Name, i.Quantity, int(i.LineTotal()))
	}
	s += fmt.Sprintf("\nTotal: ₩%d\n\n[ Submit order ]  (Enter)\n[ Keep shopping ]  (a)\n\nUp/Down=select  -=decrease  Backspace=remove  Enter=submit order  a=keep shopping", int(c.Total()))
	return s
}
