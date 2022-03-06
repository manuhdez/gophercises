package game

import (
	"fmt"
	"github.com/manuhdez/gophercises/internal/quiz/domain/question"
	"os"
	"time"
)

// Game represents a quiz game. It has a list of questions and a score.
type Game struct {
	repo      question.Repository
	questions []question.Question
	score     int
	answers   []string
	duration  time.Duration
}

// NewGame creates a new game instance
func NewGame(repo question.Repository, duration time.Duration) *Game {
	return &Game{
		repo:     repo,
		answers:  []string{},
		duration: duration,
	}
}

// Setup prepares the game to be played.
// If there is an error it exits the process with an error message.
func (g *Game) Setup() {
	questions, err := g.repo.GetAll()
	if err != nil {
		fmt.Println("Could not load questions:", err)
		os.Exit(1)
	}
	g.questions = questions
}

// Start starts the timer and the question loop.
// Once the timer or the questions are over, it displays the score and exits.
func (g Game) Start() {
	g.StartTimer()
	g.AskQuestions()
	g.PrintScore()
}

// StartTimer starts a timer for the game
func (g *Game) StartTimer() {
	timer := time.NewTimer(g.duration)

	go func() {
		<-timer.C
		fmt.Println("\nTime is up!")
		g.PrintScore()
	}()

}

// AskQuestions asks the questions in a loop.
func (g *Game) AskQuestions() {
	for i, q := range g.questions {
		fmt.Printf("Title %d: %s\n", i+1, q.Title)

		// Reads the answer from the user
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			return
		}

		g.CheckAnswer(q, answer)
	}
}

// CheckAnswer checks if the answer is correct and updates the score.
func (g *Game) CheckAnswer(question question.Question, answer string) {
	if question.IsCorrect(answer) {
		g.score++
	}
}

// PrintScore prints the score of the game and exits.
func (g *Game) PrintScore() {
	fmt.Printf("Your final score is %d out of %d\n", g.score, len(g.questions))
	os.Exit(0)
}
