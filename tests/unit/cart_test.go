package unit

import (
	"testing"

	"shop.gyeongho.dev/internal/model"
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

// M1: RemoveItem removes item at index and updates Items; out of range is no-op.
func TestCartRemoveItem(t *testing.T) {
	c := &model.Cart{
		Items: []model.CartItem{
			{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 2},
			{ProductID: "2", Name: "B", UnitPrice: 5, Quantity: 1},
		},
	}
	c.RemoveItem(1)
	if len(c.Items) != 1 {
		t.Fatalf("expected 1 item after RemoveItem(1), got %d", len(c.Items))
	}
	if c.Items[0].ProductID != "1" || c.Items[0].Quantity != 2 {
		t.Errorf("remaining item should be A x2, got %s x%d", c.Items[0].Name, c.Items[0].Quantity)
	}
	c.RemoveItem(0)
	if len(c.Items) != 0 {
		t.Errorf("expected 0 items after RemoveItem(0), got %d", len(c.Items))
	}
	// Out of range no-op
	c2 := &model.Cart{Items: []model.CartItem{{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1}}}
	c2.RemoveItem(-1)
	c2.RemoveItem(1)
	if len(c2.Items) != 1 {
		t.Errorf("out-of-range RemoveItem should not change items, got len %d", len(c2.Items))
	}
}

// M2: DecreaseQuantity decreases item at index by 1; if quantity was 1, removes item and returns true; else returns false. Out of range returns false.
func TestCartDecreaseQuantity(t *testing.T) {
	c := &model.Cart{
		Items: []model.CartItem{
			{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 3},
		},
	}
	removed := c.DecreaseQuantity(0)
	if removed {
		t.Error("DecreaseQuantity from 3 should not remove, removed=true")
	}
	if c.Items[0].Quantity != 2 {
		t.Errorf("expected quantity 2, got %d", c.Items[0].Quantity)
	}
	removed = c.DecreaseQuantity(0)
	if removed {
		t.Error("DecreaseQuantity from 2 should not remove, removed=true")
	}
	if c.Items[0].Quantity != 1 {
		t.Errorf("expected quantity 1, got %d", c.Items[0].Quantity)
	}
	removed = c.DecreaseQuantity(0)
	if !removed {
		t.Error("DecreaseQuantity from 1 should remove and return true")
	}
	if len(c.Items) != 0 {
		t.Errorf("expected 0 items after decrease to 0, got %d", len(c.Items))
	}
	// Out of range
	if c.DecreaseQuantity(-1) || c.DecreaseQuantity(0) {
		t.Error("DecreaseQuantity out of range should return false")
	}
}

// CM2: AddOrMergeItem merges by productID; empty cart then same product twice -> one item, quantity summed.
func TestCartAddOrMergeItemMergesSameProduct(t *testing.T) {
	c := &model.Cart{}
	c.AddOrMergeItem("1", "A", 10, 2, 100)
	if len(c.Items) != 1 || c.Items[0].ProductID != "1" || c.Items[0].Quantity != 2 {
		qty := 0
		if len(c.Items) > 0 {
			qty = c.Items[0].Quantity
		}
		t.Fatalf("after first AddOrMergeItem expected 1 item qty 2, got %d items qty %d", len(c.Items), qty)
	}
	c.AddOrMergeItem("1", "A", 10, 3, 100)
	if len(c.Items) != 1 || c.Items[0].Quantity != 5 {
		qty := 0
		if len(c.Items) > 0 {
			qty = c.Items[0].Quantity
		}
		t.Errorf("after second AddOrMergeItem same product expected 1 item Quantity 5, got %d items Quantity %d", len(c.Items), qty)
	}
}

// CM3: AddOrMergeItem caps quantity at maxStock.
func TestCartAddOrMergeItemCapsAtMaxStock(t *testing.T) {
	c := &model.Cart{}
	c.AddOrMergeItem("1", "A", 10, 2, 3)
	if len(c.Items) != 1 || c.Items[0].Quantity != 2 {
		t.Fatalf("expected 1 item qty 2, got %d items", len(c.Items))
	}
	c.AddOrMergeItem("1", "A", 10, 5, 3)
	if len(c.Items) != 1 || c.Items[0].Quantity != 3 {
		qty := 0
		if len(c.Items) > 0 {
			qty = c.Items[0].Quantity
		}
		t.Errorf("expected merged quantity capped at 3, got %d", qty)
	}
}
