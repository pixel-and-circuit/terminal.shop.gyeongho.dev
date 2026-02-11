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

// RemoveItem removes the item at index i. Does nothing if i is out of range.
func (c *Cart) RemoveItem(i int) {
	if i < 0 || i >= len(c.Items) {
		return
	}
	c.Items = append(c.Items[:i], c.Items[i+1:]...)
}

// DecreaseQuantity decreases the quantity of the item at index i by 1. If quantity was 1, the item is removed and true is returned; otherwise false. Returns false if i is out of range.
func (c *Cart) DecreaseQuantity(i int) (removed bool) {
	if i < 0 || i >= len(c.Items) {
		return false
	}
	if c.Items[i].Quantity <= 1 {
		c.RemoveItem(i)
		return true
	}
	c.Items[i].Quantity--
	return false
}

// AddOrMergeItem adds addQty to an existing line with the same productID, or appends a new line. Quantity is capped at maxStock.
func (c *Cart) AddOrMergeItem(productID, name string, unitPrice float64, addQty, maxStock int) {
	if addQty <= 0 {
		return
	}
	for i := range c.Items {
		if c.Items[i].ProductID == productID {
			c.Items[i].Quantity += addQty
			if c.Items[i].Quantity > maxStock {
				c.Items[i].Quantity = maxStock
			}
			return
		}
	}
	qty := addQty
	if qty > maxStock {
		qty = maxStock
	}
	c.Items = append(c.Items, CartItem{
		ProductID: productID,
		Name:      name,
		UnitPrice: unitPrice,
		Quantity:  qty,
	})
}
