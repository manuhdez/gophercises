package question

// Repository represents the contract for a question repository.
type Repository interface {
	GetAll() ([]Question, error)
}

// Question represents a question.
type Question struct {
	Title  string
	Answer string
}

// NewQuestion creates a new question.
func NewQuestion(question, answer string) Question {
	return Question{
		Title:  question,
		Answer: answer,
	}
}

// IsCorrect checks if the given answer is correct.
func (q *Question) IsCorrect(answer string) bool {
	return q.Answer == answer
}
