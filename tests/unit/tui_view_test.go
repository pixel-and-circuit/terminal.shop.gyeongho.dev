package unit

import (
	"strings"
	"testing"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/tui"
)

func TestViewContainsShopAboutFaqLabels(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	view := m.View()
	// Header shows "a shop", "s about", "d faq" (terminal.shop style)
	for _, sub := range []string{"shop", "about", "faq"} {
		if !strings.Contains(strings.ToLower(view), sub) {
			t.Errorf("View should contain %q, got:\n%s", sub, view)
		}
	}
	for _, key := range []string{"a", "s", "d"} {
		if !strings.Contains(view, key) {
			t.Errorf("View should contain shortcut %q, got:\n%s", key, view)
		}
	}
}
