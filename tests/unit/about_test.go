package unit

import (
	"strings"
	"testing"

	"shop.gyeongho.dev/internal/model"
	"shop.gyeongho.dev/internal/tui/pages"
)

func TestAboutViewShowsStoreInfo(t *testing.T) {
	info := model.StoreInfo{Title: "About", Body: "Store info here."}
	view := pages.About(info)
	if !strings.Contains(view, "About") || !strings.Contains(view, "Store info") {
		t.Errorf("About view should show title and body, got:\n%s", view)
	}
}
