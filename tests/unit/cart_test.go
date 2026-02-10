package unit

import (
	"testing"

	"mushroom.gyeongho.dev/internal/model"
)

func TestCartTotalAndItemCount(t *testing.T) {
	c := model.Cart{
		Items: []model.CartItem{
			{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 2},
			{ProductID: "2", Name: "B", UnitPrice: 5, Quantity: 1},
		},
	}
	if c.Total() != 25 {
		t.Errorf("Total expected 25, got %.2f", c.Total())
	}
	if c.ItemCount() != 3 {
		t.Errorf("ItemCount expected 3, got %d", c.ItemCount())
	}
}
