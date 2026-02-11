package pages

import "mushroom.gyeongho.dev/internal/model"

// About renders the store about page.
func About(info model.StoreInfo) string {
	if info.Title == "" && info.Body == "" {
		return "About the store.\n\nPress a/s/d to navigate."
	}
	return info.Title + "\n\n" + info.Body
}
