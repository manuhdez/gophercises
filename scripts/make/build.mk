.PHONY: build

build: build-quiz

build-quiz:
	go build -o build/quiz cmd/quiz/main.go
