package integration

import (
	"context"
	"strings"
	"testing"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

// I1: Shop add one product qty 3 → open Cart → press - twice → quantity 1, total and line total correct.
func TestCartDecreaseQuantityIntegration(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
	m.Loading = false
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	m.AddQuantity = 3
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	if len(mod.Cart.Items) != 1 || mod.Cart.Items[0].Quantity != 3 {
		t.Fatalf("expected 1 cart item with Quantity 3, got %d items", len(mod.Cart.Items))
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod = updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod = updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 item after two -, got %d", len(mod.Cart.Items))
	}
	if mod.Cart.Items[0].Quantity != 1 {
		t.Errorf("expected Quantity 1, got %d", mod.Cart.Items[0].Quantity)
	}
	expectedTotal := 800.0
	if mod.Cart.Total() != expectedTotal {
		t.Errorf("expected Total %.2f, got %.2f", expectedTotal, mod.Cart.Total())
	}
	view := mod.View()
	if !strings.Contains(view, "x1") || !strings.Contains(view, "800") {
		t.Errorf("cart view should show x1 and line total 800, got:\n%s", view)
	}
}

// I2: Same as I1 setup; then - three times → cart empty.
func TestCartDecreaseToEmptyIntegration(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
	m.Loading = false
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	m.AddQuantity = 3
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod = updated.(tui.Model)
	for i := 0; i < 3; i++ {
		updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'-'}})
		mod = updated.(tui.Model)
	}
	if len(mod.Cart.Items) != 0 {
		t.Errorf("expected empty cart after three -, got %d items", len(mod.Cart.Items))
	}
	view := mod.View()
	if !strings.Contains(strings.ToLower(view), "empty") {
		t.Errorf("cart view should show empty message, got:\n%s", view)
	}
}

// I3: Shop add two products → open Cart → Down to select second → Backspace → only first remains.
func TestCartBackspaceRemovesSelectedIntegration(t *testing.T) {
	client := apiclient.MockClient{}
	prods, _ := client.GetProducts(context.Background())
	m := tui.NewModel(client)
	m.Loading = false
	m.Products = prods
	m.CurrentPage = tui.PageShop
	m.Cursor = 0
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod := updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod = updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyEnter})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 2 {
		t.Fatalf("expected 2 cart items, got %d", len(mod.Cart.Items))
	}
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'c'}})
	mod = updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyDown})
	mod = updated.(tui.Model)
	updated, _ = mod.Update(tea.KeyMsg{Type: tea.KeyBackspace})
	mod = updated.(tui.Model)
	if len(mod.Cart.Items) != 1 {
		t.Fatalf("expected 1 item after Backspace, got %d", len(mod.Cart.Items))
	}
	if mod.Cart.Items[0].Name != "Oyster Mushroom" {
		t.Errorf("expected remaining item Oyster Mushroom, got %s", mod.Cart.Items[0].Name)
	}
	view := mod.View()
	if !strings.Contains(view, "Oyster Mushroom") || strings.Contains(view, "Enoki") {
		t.Errorf("cart view should show only Oyster Mushroom, got:\n%s", view)
	}
}
