package main

import (
	"fmt"
	"github.com/manuhdez/gophercises/internal/shortener/handlers"
	"github.com/manuhdez/gophercises/internal/shortener/src/application"
	"github.com/manuhdez/gophercises/internal/shortener/src/infra"
	"log"
	"net/http"
	"os"
)

func main() {
	repository := infra.NewYamlRedirectRepository()
	finder := application.NewRedirectionFinder(repository)

	mux := http.NewServeMux()
	mux.Handle("/status", handlers.GetStatusHandler())
	mux.Handle("/", handlers.GetRedirectionHandler(finder))

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	fmt.Printf("Server listening on port %s\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux))
}
