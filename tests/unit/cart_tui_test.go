package unit

import (
	"strings"
	"testing"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/model"
	"shop.gyeongho.dev/internal/tui"
	"shop.gyeongho.dev/internal/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
)

// U5: Cart entry (KeyCart) sets cursor to 0.
func TestCartEntrySetsCursorZero(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageShop
	m.Cursor = 3
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod := updated.(tui.Model)
	if mod.CurrentPage != tui.PageCart {
		t.Fatalf("expected PageCart, got %v", mod.CurrentPage)
	}
	if mod.Cursor != 0 {
		t.Errorf("expected Cursor 0 on Cart entry, got %d", mod.Cursor)
	}
}

// U1: On Cart, KeyDown moves cursor to next item (or keeps at last). Empty cart no-op.
func TestCartKeyDownMovesCursor(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cursor = 0
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
		{ProductID: "2", Name: "B", UnitPrice: 5, Quantity: 1},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod := updated.(tui.Model)
	if mod.Cursor != 1 {
		t.Errorf("expected Cursor 1 after KeyDown, got %d", mod.Cursor)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod = updated.(tui.Model)
	if mod.Cursor != 1 {
		t.Errorf("expected Cursor still 1 at last item, got %d", mod.Cursor)
	}
	// Empty cart: KeyDown no-op
	m2 := tui.NewModel(apiclient.MockClient{})
	m2.CurrentPage = tui.PageCart
	m2.Cart.Items = nil
	updated, _ = m2.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod = updated.(tui.Model)
	if mod.Cursor != 0 {
		t.Errorf("empty cart KeyDown should leave Cursor 0, got %d", mod.Cursor)
	}
}

// U2: On Cart, KeyUp moves cursor to previous item (or keeps at 0). Empty cart no-op.
func TestCartKeyUpMovesCursor(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cursor = 1
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
		{ProductID: "2", Name: "B", UnitPrice: 5, Quantity: 1},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyUp})
	mod := updated.(tui.Model)
	if mod.Cursor != 0 {
		t.Errorf("expected Cursor 0 after KeyUp, got %d", mod.Cursor)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyUp})
	mod = updated.(tui.Model)
	if mod.Cursor != 0 {
		t.Errorf("expected Cursor still 0 at first item, got %d", mod.Cursor)
	}
}

// U3: On Cart, KeyMinus decreases selected item; if quantity was 1, item is removed. Empty cart no-op.
func TestCartKeyMinusDecreasesOrRemoves(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cursor = 0
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 2},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 1 || mod.Cart.Items[0].Quantity != 1 {
		qty := 0
		if len(mod.Cart.Items) > 0 {
			qty = mod.Cart.Items[0].Quantity
		}
		t.Errorf("expected one item with Quantity 1, got %d items and qty %d", len(mod.Cart.Items), qty)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 0 {
		t.Errorf("expected cart empty after second -, got %d items", len(mod.Cart.Items))
	}
	if mod.Cursor != 0 {
		t.Errorf("expected Cursor 0 when cart empty, got %d", mod.Cursor)
	}
}

// U4: On Cart, Backspace removes selected item. Empty cart no-op.
func TestCartBackspaceRemovesItem(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cursor = 1
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
		{ProductID: "2", Name: "B", UnitPrice: 5, Quantity: 2},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyBackspace})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 item after Backspace, got %d", len(mod.Cart.Items))
	}
	if mod.Cart.Items[0].ProductID != "1" {
		t.Errorf("expected remaining item A, got %s", mod.Cart.Items[0].Name)
	}
	if mod.Cursor != 0 {
		t.Errorf("cursor should clamp to 0 after removing second item, got %d", mod.Cursor)
	}
}

// U6: After removing last item, cursor 0; KeyDown/KeyUp no-op on empty cart.
func TestCartEmptyAfterRemoveKeyUpDownNoOp(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cursor = 0
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyBackspace})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 0 || mod.Cursor != 0 {
		t.Fatalf("expected empty cart and Cursor 0, got %d items cursor %d", len(mod.Cart.Items), mod.Cursor)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod = updated.(tui.Model)
	if mod.Cursor != 0 {
		t.Errorf("KeyDown on empty cart should leave Cursor 0, got %d", mod.Cursor)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyUp})
	mod = updated.(tui.Model)
	if mod.Cursor != 0 {
		t.Errorf("KeyUp on empty cart should leave Cursor 0, got %d", mod.Cursor)
	}
}

// U7: On Shop/About/FAQ/Landing, KeyMinus and Backspace do not change Cart.
func TestCartKeyMinusBackspaceOffCartNoEffect(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 2},
	}
	// About page: KeyMinus should not change cart
	m.CurrentPage = tui.PageAbout
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 1 || mod.Cart.Items[0].Quantity != 2 {
		qty := 0
		if len(mod.Cart.Items) > 0 {
			qty = mod.Cart.Items[0].Quantity
		}
		t.Errorf("on About, KeyMinus should not change cart, got %d items qty %d", len(mod.Cart.Items), qty)
	}
	// About page: Backspace should not change cart
	m.CurrentPage = tui.PageAbout
	m.Cart.Items = []model.CartItem{{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1}}
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyBackspace})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Errorf("on About, Backspace should not remove cart item, got %d items", len(mod.Cart.Items))
	}
}

// V1: Cart with items shows cursor (e.g. ">") on selected index row.
func TestCartViewShowsCursorOnSelectedItem(t *testing.T) {
	c := model.Cart{
		Items: []model.CartItem{
			{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
			{ProductID: "2", Name: "B", UnitPrice: 5, Quantity: 2},
		},
	}
	view := pages.Cart(c, 1)
	if !strings.Contains(view, ">") {
		t.Errorf("Cart view with selected index 1 should show cursor >, got:\n%s", view)
	}
	// Selected row (B) should have the cursor
	lines := strings.Split(view, "\n")
	foundSelected := false
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), ">") && strings.Contains(line, "B") {
			foundSelected = true
			break
		}
	}
	if !foundSelected {
		t.Errorf("Cart view should show > on selected item line, got:\n%s", view)
	}
}

// V2: Empty cart view has no cursor/selection marker.
func TestCartViewEmptyNoCursor(t *testing.T) {
	c := model.Cart{Items: nil}
	view := pages.Cart(c, 0)
	if strings.Contains(view, ">") {
		t.Errorf("Empty cart view should not show cursor, got:\n%s", view)
	}
	if !strings.Contains(view, "empty") {
		t.Errorf("Empty cart view should contain empty message, got:\n%s", view)
	}
}

// U8: On Cart with items, Enter opens order modal (Fake Door).
func TestCartEnterWithItemsOpensOrderModal(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cart.Items = []model.CartItem{
		{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1},
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	if !mod.OrderModalOpen {
		t.Errorf("expected OrderModalOpen true after Enter on cart with items, got false")
	}
}

// U9: On Cart when empty, Enter does not open modal.
func TestCartEnterEmptyCartNoModal(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cart.Items = nil
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	if mod.OrderModalOpen {
		t.Errorf("expected OrderModalOpen false when cart empty and Enter, got true")
	}
}

// U10: When order modal is open, Enter or Esc closes it.
func TestOrderModalCloseWithEnterOrEscape(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.OrderModalOpen = true
	m.Cart.Items = []model.CartItem{{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1}}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	if mod.OrderModalOpen {
		t.Errorf("expected OrderModalOpen false after Enter in modal, got true")
	}
	m.OrderModalOpen = true
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyEscape})
	mod = updated.(tui.Model)
	if mod.OrderModalOpen {
		t.Errorf("expected OrderModalOpen false after Esc in modal, got true")
	}
}

// U11: On Cart, Keep shopping (a) navigates to Shop.
func TestCartKeepShoppingGoesToShop(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageCart
	m.Cart.Items = []model.CartItem{{ProductID: "1", Name: "A", UnitPrice: 10, Quantity: 1}}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})
	mod := updated.(tui.Model)
	if mod.CurrentPage != tui.PageShop {
		t.Errorf("expected PageShop after a on Cart (keep shopping), got %v", mod.CurrentPage)
	}
}
