package unit

import (
	"context"
	"strings"
	"testing"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/tui"
	"mushroom.gyeongho.dev/internal/tui/pages"

	tea "github.com/charmbracelet/bubbletea"
)

func TestShopViewShowsProductList(t *testing.T) {
	prods, _ := (apiclient.MockClient{}).GetProducts(context.Background())
	view := pages.Shop(prods, 0, 0, 60, 1)
	if view == "" {
		t.Fatal("Shop view should not be empty when products exist")
	}
	if !strings.Contains(view, "Oyster Mushroom") {
		t.Errorf("Shop view should contain product name Oyster Mushroom, got:\n%s", view)
	}
}

func TestShopUpdateUpDownChangesCursor(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	prods, _ := (apiclient.MockClient{}).GetProducts(context.Background())
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	// Down increases cursor
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod := updated.(tui.Model)
	if mod.Cursor != 1 {
		t.Errorf("expected Cursor 1 after Down, got %d", mod.Cursor)
	}
	// Up decreases cursor
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyUp})
	mod = updated.(tui.Model)
	if mod.Cursor != 0 {
		t.Errorf("expected Cursor 0 after Up, got %d", mod.Cursor)
	}
}

// TestUpdateKeyPlusOnShopIncreasesAddQuantity: On Shop page, + increases AddQuantity.
func TestUpdateKeyPlusOnShopIncreasesAddQuantity(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	prods, _ := (apiclient.MockClient{}).GetProducts(context.Background())
	m.Products = prods
	// Default AddQuantity is 1; after + should be 2
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'+'}})
	mod := updated.(tui.Model)
	if mod.AddQuantity != 2 {
		t.Errorf("expected AddQuantity 2 after first +, got %d", mod.AddQuantity)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'+'}})
	mod = updated.(tui.Model)
	if mod.AddQuantity != 3 {
		t.Errorf("expected AddQuantity 3 after second +, got %d", mod.AddQuantity)
	}
}

// TestUpdateKeyMinusOnShopDecreasesAddQuantity: On Shop page, - decreases AddQuantity (min 1).
func TestUpdateKeyMinusOnShopDecreasesAddQuantity(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	m.AddQuantity = 3
	// - twice -> 2, then 1
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod := updated.(tui.Model)
	if mod.AddQuantity != 2 {
		t.Errorf("expected AddQuantity 2 after first -, got %d", mod.AddQuantity)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod = updated.(tui.Model)
	if mod.AddQuantity != 1 {
		t.Errorf("expected AddQuantity 1 after second -, got %d", mod.AddQuantity)
	}
	// - at 1 should stay 1
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod = updated.(tui.Model)
	if mod.AddQuantity != 1 {
		t.Errorf("expected AddQuantity still 1 after - at 1, got %d", mod.AddQuantity)
	}
}

// TestUpdateEnterAddsAddQuantityToCart: On Shop, Enter adds AddQuantity to cart (one line with that quantity).
func TestUpdateEnterAddsAddQuantityToCart(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	prods, _ := (apiclient.MockClient{}).GetProducts(context.Background())
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	m.AddQuantity = 3
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 cart item after Enter, got %d", len(mod.Cart.Items))
	}
	if mod.Cart.Items[0].Quantity != 3 {
		t.Errorf("expected cart item Quantity 3, got %d", mod.Cart.Items[0].Quantity)
	}
}

// T1: Same product Enter 3 times (cursor unchanged) -> one cart row with Quantity 3 (merge by productID).
func TestUpdateEnterSameItemMultipleTimesMergesQuantity(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	prods, _ := (apiclient.MockClient{}).GetProducts(context.Background())
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	m.AddQuantity = 1
	for i := 0; i < 3; i++ {
		updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
		m = updated.(tui.Model)
	}
	if len(m.Cart.Items) != 1 {
		t.Fatalf("expected 1 cart item after 3 Enter on same product, got %d", len(m.Cart.Items))
	}
	if m.Cart.Items[0].Quantity != 3 {
		t.Errorf("expected merged Quantity 3, got %d", m.Cart.Items[0].Quantity)
	}
}

// TestUpdateKeyPlusMinusOffShopNoEffect: +/- on non-Shop page do not change AddQuantity (or are no-op).
func TestUpdateKeyPlusMinusOffShopNoEffect(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageAbout
	m.AddQuantity = 1
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'+'}})
	mod := updated.(tui.Model)
	if mod.AddQuantity != 1 {
		t.Errorf("expected AddQuantity unchanged 1 on About after +, got %d", mod.AddQuantity)
	}
}
