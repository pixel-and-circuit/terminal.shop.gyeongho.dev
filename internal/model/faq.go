package model

// FAQEntry is one Q&A pair for the FAQ page.
type FAQEntry struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
