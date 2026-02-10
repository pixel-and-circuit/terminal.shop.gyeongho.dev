package pages

import (
	"fmt"
	"mushroom.gyeongho.dev/internal/model"
)

// FAQ renders the FAQ page with scroll.
func FAQ(entries []model.FAQEntry, scrollOffset, width int) string {
	if len(entries) == 0 {
		return "No FAQ. Press a/s/d to navigate."
	}
	var s string
	for i, e := range entries {
		if i < scrollOffset {
			continue
		}
		s += fmt.Sprintf("Q: %s\nA: %s\n\n", e.Question, e.Answer)
	}
	return s + "Up/Down=scroll  a/s/d=nav"
}
