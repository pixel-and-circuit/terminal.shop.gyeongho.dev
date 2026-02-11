package testserver

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// NewFixtureHandler returns an http.Handler that serves API responses by reading
// JSON files from fixtureDir. Used for E2E tests; developers edit the JSON
// files to change mock responses without code changes.
//
// Mapping: GET /products -> products.json, GET /about -> about.json,
// GET /faq -> faq.json, GET /cart and POST /cart -> cart.json,
// POST /orders -> order.json.
// If the request has X-User-Id, the handler echoes it in response header X-Echo-User-Id.
func NewFixtureHandler(fixtureDir string) http.Handler {
	return &fixtureHandler{dir: fixtureDir}
}

type fixtureHandler struct {
	dir string
}

func (h *fixtureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if userID := r.Header.Get("X-User-Id"); userID != "" {
		w.Header().Set("X-Echo-User-Id", userID)
	}
	w.Header().Set("Content-Type", "application/json")

	var filename string
	statusCreated := false
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/products":
		filename = "products.json"
	case r.Method == http.MethodGet && r.URL.Path == "/about":
		filename = "about.json"
	case r.Method == http.MethodGet && r.URL.Path == "/faq":
		filename = "faq.json"
	case r.URL.Path == "/cart":
		filename = "cart.json"
	case r.Method == http.MethodPost && r.URL.Path == "/orders":
		filename = "order.json"
		statusCreated = true
	default:
		http.NotFound(w, r)
		return
	}

	path := filepath.Join(h.dir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("testserver: read %s: %v", path, err)
		http.Error(w, "fixture read error", http.StatusInternalServerError)
		return
	}
	if statusCreated {
		w.WriteHeader(http.StatusCreated)
	}
	if _, err := w.Write(data); err != nil {
		log.Printf("testserver: write: %v", err)
	}
}
