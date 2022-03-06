package main

import (
	"fmt"
	"github.com/manuhdez/gophercises/internal/shortener/handlers"
	"github.com/manuhdez/gophercises/internal/shortener/src/application"
	"github.com/manuhdez/gophercises/internal/shortener/src/infra"
	"log"
	"net/http"
)

func main() {
	repository := infra.NewRedirectionRepository()
	finder := application.NewRedirectionFinder(repository)

	mux := http.NewServeMux()
	mux.Handle("/status", handlers.GetStatusHandler())
	mux.Handle("/", handlers.GetRedirectionHandler(finder))

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
