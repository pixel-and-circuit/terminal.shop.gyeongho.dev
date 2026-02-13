package unit

import (
	"testing"

	"shop.gyeongho.dev/internal/apiclient"
	"shop.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdateKeyASetsPageToShop(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageLanding
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
	updated, _ := m.Update(msg)
	model := updated.(tui.Model)
	if model.CurrentPage != tui.PageShop {
		t.Errorf("expected CurrentPage Shop after 'a', got %v", model.CurrentPage)
	}
}

func TestUpdateKeySSetsPageToAbout(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageLanding
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'s'}}
	updated, _ := m.Update(msg)
	model := updated.(tui.Model)
	if model.CurrentPage != tui.PageAbout {
		t.Errorf("expected CurrentPage About after 's', got %v", model.CurrentPage)
	}
}

func TestUpdateKeyDSetsPageToFAQ(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.CurrentPage = tui.PageLanding
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}}
	updated, _ := m.Update(msg)
	model := updated.(tui.Model)
	if model.CurrentPage != tui.PageFAQ {
		t.Errorf("expected CurrentPage FAQ after 'd', got %v", model.CurrentPage)
	}
}
