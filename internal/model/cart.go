package model

import "time"

// CartItem is one line in the cart: product snapshot + quantity.
type CartItem struct {
	ProductID string  `json:"productId"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unitPrice"`
	Quantity  int     `json:"quantity"`
}

// LineTotal returns UnitPrice * Quantity.
func (c CartItem) LineTotal() float64 {
	return c.UnitPrice * float64(c.Quantity)
}

// Cart is the current shopping cart state.
type Cart struct {
	Items     []CartItem `json:"items"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
}

// Total returns the sum of line totals.
func (c Cart) Total() float64 {
	var t float64
	for _, i := range c.Items {
		t += i.LineTotal()
	}
	return t
}

// ItemCount returns total number of units.
func (c Cart) ItemCount() int {
	var n int
	for _, i := range c.Items {
		n += i.Quantity
	}
	return n
}
