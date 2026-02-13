package integration

import (
	"testing"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

// TestNavigationKeyASwitchesToShop simulates starting program and sending "a".
func TestNavigationKeyASwitchesToShop(t *testing.T) {
	client := apiclient.MockClient{}
	m := tui.NewModel(client)
	// Simulate key "a"
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}})
	if cmd != nil {
		t.Log("cmd ignored for this test")
	}
	model, ok := updated.(tui.Model)
	if !ok {
		t.Fatal("Update should return tui.Model")
	}
	if model.CurrentPage != tui.PageShop {
		t.Errorf("after key 'a', current page should be Shop, got %v", model.CurrentPage)
	}
}
