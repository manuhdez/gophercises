package infra

import (
	"fmt"
	"github.com/manuhdez/gophercises/internal/shortener/src/domain"
	"gopkg.in/yaml.v2"
	"os"
)

type YamlData struct {
	Redirects []struct {
		From string `yaml:"from"`
		To   string `yaml:"to"`
	} `yaml:"redirects"`
}

type YamlRedirectRepository struct {
}

func NewYamlRedirectRepository() YamlRedirectRepository {
	return YamlRedirectRepository{}
}

func (YamlRedirectRepository) FindByShortcode(shortcode string) (domain.Redirection, error) {
	redirectsMap, err := readYamlFile()
	if err != nil {
		return domain.Redirection{}, err
	}

	redirect, ok := redirectsMap[shortcode]
	if !ok {
		return domain.Redirection{}, fmt.Errorf("redirection not found")
	}

	return redirect, nil
}

func readYamlFile() (map[string]domain.Redirection, error) {
	file, err := os.ReadFile("data/redirects.yaml")
	if err != nil {
		fmt.Println("error opening yaml file: ", err)
		return make(map[string]domain.Redirection), err
	}

	// parse yaml file
	var t YamlData
	err = yaml.Unmarshal(file, &t)
	if err != nil {
		fmt.Printf("error reading yaml file: %e", err)
		return make(map[string]domain.Redirection), err
	}

	return parseRedirectsToMap(t), nil
}

func parseRedirectsToMap(redirects YamlData) map[string]domain.Redirection {
	redirectMap := make(map[string]domain.Redirection)
	for _, r := range redirects.Redirects {
		redirectMap[r.From] = domain.NewRedirection(r.From, r.To)
	}
	return redirectMap
}
