package pages

import (
	"fmt"
	"mushroom.gyeongho.dev/internal/model"
)

// FAQ renders the full FAQ page. Scrolling is handled by the app viewport.
func FAQ(entries []model.FAQEntry, width int) string {
	if len(entries) == 0 {
		return "No FAQ. Press a/s/d to navigate."
	}
	var s string
	for _, e := range entries {
		s += fmt.Sprintf("Q: %s\nA: %s\n\n", e.Question, e.Answer)
	}
	return s + "Up/Down=scroll  a/s/d=nav"
}
