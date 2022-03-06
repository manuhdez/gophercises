package main

import (
	"flag"
	"github.com/manuhdez/gophercises/internal/quiz/domain/game"
	"github.com/manuhdez/gophercises/internal/quiz/infra"
	"time"
)

const (
	defaultDuration = 30
	defaultFilename = "problems.csv"
)

func main() {
	filename, duration := readFlags()
	repo := infra.NewQuestionRepository(filename)

	g := game.NewGame(repo, duration)
	g.Setup()
	g.Start()
}

func readFlags() (string, time.Duration) {
	filename := flag.String("file", defaultFilename, "file with questions")
	duration := flag.Int("time", defaultDuration, "Set a custom game duration")
	flag.Parse()

	return *filename, time.Second * time.Duration(*duration)
}
