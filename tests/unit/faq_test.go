package unit

import (
	"strings"
	"testing"

	"shop.gyeongho.dev/internal/model"
	"shop.gyeongho.dev/internal/tui/pages"
)

func TestFAQViewShowsQuestionsAndAnswers(t *testing.T) {
	entries := []model.FAQEntry{
		{Question: "Q1?", Answer: "A1"},
	}
	view := pages.FAQ(entries, 60)
	if !strings.Contains(view, "Q1?") || !strings.Contains(view, "A1") {
		t.Errorf("FAQ view should show Q&A, got:\n%s", view)
	}
}
