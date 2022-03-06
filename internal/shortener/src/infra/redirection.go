package infra

import (
	"errors"
	"github.com/manuhdez/gophercises/internal/shortener/src/domain"
)

type RedirectionRepository struct {
	redirections map[string]string
}

func NewRedirectionRepository() RedirectionRepository {
	mapHandler := map[string]string{
		"g": "https://google.com",
		"d": "https://dev.to",
		"t": "https://twitter.com",
	}

	return RedirectionRepository{
		redirections: mapHandler,
	}
}

func (r RedirectionRepository) FindByShortcode(short string) (domain.Redirection, error) {
	uri, ok := r.redirections[short]
	if !ok {
		return domain.Redirection{}, errors.New("redirection not found")
	}

	return domain.NewRedirection(short, uri), nil
}
