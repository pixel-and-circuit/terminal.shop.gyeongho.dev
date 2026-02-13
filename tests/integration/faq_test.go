package integration

import (
	"context"
	"strings"
	"testing"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func TestPressDShowsFAQScroll(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false
	faq, _ := (apiclient.MockClient{}).GetFAQ(context.Background())
	m.FAQ = faq
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}})
	mod := updated.(tui.Model)
	mod.Loading = false
	if mod.CurrentPage != tui.PageFAQ {
		t.Fatalf("expected PageFAQ after d, got %v", mod.CurrentPage)
	}
	view := mod.View()
	if !strings.Contains(view, "order") && !strings.Contains(view, "payment") {
		t.Errorf("faq view should show FAQ content, got:\n%s", view)
	}
}
