package integration

import (
	"context"
	"strings"
	"testing"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func TestOpenShopAddToCartOpenCart(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
	m.Loading = false
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	// Add to cart (Enter)
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 cart item after Enter, got %d", len(mod.Cart.Items))
	}
	// Open cart (c)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod = updated.(tui.Model)
	mod.Loading = false
	if mod.CurrentPage != tui.PageCart {
		t.Errorf("expected PageCart after c, got %v", mod.CurrentPage)
	}
	view := mod.View()
	if !strings.Contains(view, "Oyster") || !strings.Contains(view, "800") {
		t.Errorf("cart view should show product and price, got:\n%s", view)
	}
}

func TestShopPlusMinusQtyThenEnterAddsThatQuantityToCart(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
	m.Loading = false
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	// Increase add quantity with + twice (1 -> 2 -> 3)
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'+'}})
	mod := updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'+'}})
	mod = updated.(tui.Model)
	// Add to cart (Enter)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 cart item after Enter, got %d", len(mod.Cart.Items))
	}
	if mod.Cart.Items[0].Quantity != 3 {
		t.Errorf("expected cart item Quantity 3 (after ++), got %d", mod.Cart.Items[0].Quantity)
	}
	// Open cart and verify line total (3 * unit price)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod = updated.(tui.Model)
	mod.Loading = false
	view := mod.View()
	if !strings.Contains(view, "x3") {
		t.Errorf("cart view should show quantity x3, got:\n%s", view)
	}
}

// T2: Same product Enter 2 times -> open Cart -> view shows one row with quantity 2 (merge by productID).
func TestShopSameItemEnterMultipleTimesOneRowInCart(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
	m.Loading = false
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	m.AddQuantity = 1
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 cart item after 2 Enter on same product, got %d", len(mod.Cart.Items))
	}
	if mod.Cart.Items[0].Quantity != 2 {
		t.Errorf("expected merged Quantity 2, got %d", mod.Cart.Items[0].Quantity)
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod = updated.(tui.Model)
	mod.Loading = false
	view := mod.View()
	if !strings.Contains(view, "x2") {
		t.Errorf("cart view should show one row with x2, got:\n%s", view)
	}
	if !strings.Contains(view, "Oyster") {
		t.Errorf("cart view should show product name Oyster, got:\n%s", view)
	}
}
