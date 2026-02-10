package integration

import (
	"context"
	"strings"
	"testing"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func TestOpenShopAddToCartOpenCart(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
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
	if mod.CurrentPage != tui.PageCart {
		t.Errorf("expected PageCart after c, got %v", mod.CurrentPage)
	}
	view := mod.View()
	if !strings.Contains(view, "Shiitake") || !strings.Contains(view, "12.50") {
		t.Errorf("cart view should show product and price, got:\n%s", view)
	}
}
