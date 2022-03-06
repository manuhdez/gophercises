package domain

type RedirectionRepository interface {
	FindByShortcode(shortcode string) (Redirection, error)
}

type Redirection struct {
	Shortcode string
	URL       string
}

func NewRedirection(shortcode string, url string) Redirection {
	return Redirection{
		Shortcode: shortcode,
		URL:       url,
	}
}
