package unit

import (
	"strings"
	"testing"

	"mushroom.gyeongho.dev/internal/apiclient"
	"mushroom.gyeongho.dev/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func TestViewContainsShopAboutFaqLabels(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false // main view
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

func TestViewWhenLoadingTrue_ShowsCenteredLoaderWithInnerText(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = true
	m.Width = 80
	m.Height = 24
	view := m.View()
	if view == "" {
		t.Error("View when Loading is true should not be empty")
	}
	if !strings.Contains(view, "Loading mushroom.gyeongho.dev") {
		t.Errorf("View when Loading is true should contain loader inner text, got:\n%s", view)
	}
	// Centered output has multiple lines (viewport height)
	lines := strings.Count(view, "\n")
	if lines < 10 {
		t.Errorf("View when Loading is true should be centered (many lines), got %d lines", lines+1)
	}
}

func TestViewWhenLoadingFalse_ShowsCenteredMainViewWithHeaderAndFooter(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false
	m.Width = 80
	m.Height = 24
	view := m.View()
	if view == "" {
		t.Error("View when Loading is false should not be empty")
	}
	if !strings.Contains(strings.ToLower(view), "mushroom") {
		t.Errorf("View when Loading is false should contain header (mushroom), got:\n%s", view)
	}
	if !strings.Contains(view, "quit") {
		t.Errorf("View when Loading is false should contain footer (quit), got:\n%s", view)
	}
	lines := strings.Count(view, "\n")
	if lines < 5 {
		t.Errorf("View when Loading is false should be centered (multiple lines), got %d lines", lines+1)
	}
}

// TestViewWithSmallHeight_MenuVisibleAndFitsViewport reproduces the issue where
// reducing terminal height cuts off the menu. View() must return at most height lines
// and the header (menu) must appear in the output so it is visible when displayed.
func TestViewWithSmallHeight_MenuVisibleAndFitsViewport(t *testing.T) {
	smallHeight := 12
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false
	m.Width = 80
	m.Height = smallHeight
	view := m.View()
	assertViewFitsHeightAndMenuVisible(t, view, smallHeight, "initial small height")
}

// TestViewAfterResizeToSmallHeight_MenuVisibleAndFitsViewport simulates user resizing
// terminal from large to small: after WindowSizeMsg with small height, View() must
// fit in viewport and show menu (reproduces real-world cut-off issue).
func TestViewAfterResizeToSmallHeight_MenuVisibleAndFitsViewport(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false
	m.Width = 80
	m.Height = 30
	m.View() // one frame at large height

	// User shrinks terminal height (e.g. from 30 to 12)
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 12})
	mod, ok := updated.(tui.Model)
	if !ok {
		t.Fatalf("Update returned %T", updated)
	}
	view := mod.View()
	assertViewFitsHeightAndMenuVisible(t, view, 12, "after resize to height=12")
}

func assertViewFitsHeightAndMenuVisible(t *testing.T, view string, maxLines int, scenario string) {
	t.Helper()
	lines := strings.Split(view, "\n")
	lineCount := len(lines)
	if lineCount > maxLines {
		t.Errorf("%s: View() must return at most %d lines so menu is not cut off; got %d lines",
			scenario, maxLines, lineCount)
	}
	viewLower := strings.ToLower(view)
	if !strings.Contains(viewLower, "mushroom") || !strings.Contains(viewLower, "shop") {
		t.Errorf("%s: View() must contain menu (header)", scenario)
	}
	foundInVisible := false
	for i := 0; i < lineCount && i < 5; i++ {
		if strings.Contains(strings.ToLower(lines[i]), "mushroom") || strings.Contains(strings.ToLower(lines[i]), "shop") {
			foundInVisible = true
			break
		}
	}
	if !foundInVisible {
		t.Errorf("%s: menu should appear in first lines (height=%d); got %d lines", scenario, maxLines, lineCount)
	}
}

func TestViewAfterWindowSizeMsg_UsesNewDimensions(t *testing.T) {
	m := tui.NewModel(apiclient.MockClient{})
	m.Loading = false
	m.Width = 80
	m.Height = 24
	m.View()

	// Send resize message
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 30})
	mod, ok := updated.(tui.Model)
	if !ok {
		t.Fatalf("Update returned %T, expected tui.Model", updated)
	}
	if mod.Width != 120 || mod.Height != 30 {
		t.Errorf("model after WindowSizeMsg should have Width=120 Height=30, got %d×%d", mod.Width, mod.Height)
	}
	viewLarge := mod.View()
	// Fixed-height layout: view has header+body+footer; after resize Place uses new dimensions for centering
	if viewLarge == "" {
		t.Error("View after resize should not be empty")
	}
	if !strings.Contains(viewLarge, "mushroom") {
		t.Error("View after resize should still contain header content")
	}
}
