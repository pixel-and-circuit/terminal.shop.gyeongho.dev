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
	view := pages.Shop(prods, 0, 0, 60)
	if view == "" {
		t.Fatal("Shop view should not be empty when products exist")
	}
	if !strings.Contains(view, "Shiitake") {
		t.Errorf("Shop view should contain product name Shiitake, got:\n%s", view)
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
