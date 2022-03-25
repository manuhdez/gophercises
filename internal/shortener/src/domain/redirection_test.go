package domain

import "testing"

func TestNewRedirection(t *testing.T) {
	t.Run("should return a new redirection", func(t *testing.T) {
		redirection := NewRedirection("g", "https://www.google.com")

		if redirection.URL != "https://www.google.com" {
			t.Errorf("Expected %s, got %s", "https://www.google.com", redirection.URL)
		}
	})
}
