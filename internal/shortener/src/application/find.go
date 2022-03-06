package application

import "github.com/manuhdez/gophercises/internal/shortener/src/domain"

type RedirectionFinder struct {
	repository domain.RedirectionRepository
}

func NewRedirectionFinder(repo domain.RedirectionRepository) RedirectionFinder {
	return RedirectionFinder{repository: repo}
}

func (f *RedirectionFinder) Find(shortcode string) (domain.Redirection, error) {
	return f.repository.FindByShortcode(shortcode)
}
