package model

// FAQEntry is one Q&A pair for the FAQ page.
type FAQEntry struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
