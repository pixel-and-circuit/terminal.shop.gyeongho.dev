package integration

import (
	"context"
	"strings"
	"testing"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func TestPressSShowsAboutContent(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false
	about, _ := (apiclient.MockClient{}).GetAbout(context.Background())
	m.About = about
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}})
	mod := updated.(tui.Model)
	mod.Loading = false
	if mod.CurrentPage != tui.PageAbout {
		t.Fatalf("expected PageAbout after s, got %v", mod.CurrentPage)
	}
	view := mod.View()
	viewLower := strings.ToLower(view)
	if !strings.Contains(viewLower, "about") || !strings.Contains(viewLower, "store") {
		t.Errorf("about view should show store info, got:\n%s", view)
	}
}
