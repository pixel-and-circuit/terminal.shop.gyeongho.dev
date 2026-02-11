package pages

import (
	"fmt"
	"mushroom.gyeongho.dev/internal/model"
)

// Cart renders the cart page: items, total, and checkout hint.
func Cart(c model.Cart) string {
	if len(c.Items) == 0 {
		return "Cart is empty.\n\nPress a to go to shop."
	}
	var s string
	for _, i := range c.Items {
		s += fmt.Sprintf("  %s x%d $%.2f\n", i.Name, i.Quantity, i.LineTotal())
	}
	s += fmt.Sprintf("\nTotal: $%.2f\n\nEnter=checkout  a=shop", c.Total())
	return s
}
