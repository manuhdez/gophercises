package handlers

import (
	"github.com/manuhdez/gophercises/internal/shortener/src/application"
	"net/http"
)

// GetRedirectionHandler is a handler for GET /{shortcode}
// It receives the use case via parameter -> application.Redirect
func GetRedirectionHandler(finder application.RedirectionFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]

		redirection, err := finder.Find(path)
		if err != nil {
			http.Error(w, "Shortcode not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, redirection.URL, http.StatusMovedPermanently)
	}
}
