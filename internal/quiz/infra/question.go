package infra

import (
	"encoding/csv"
	"fmt"
	"github.com/manuhdez/gophercises/internal/quiz/domain/question"
	"os"
)

const (
	filePath = "internal/quiz/assets/"
)

type QuestionRepository struct {
	path string
}

func NewQuestionRepository(filename string) *QuestionRepository {
	return &QuestionRepository{path: filePath + filename}
}

func (r *QuestionRepository) GetAll() ([]question.Question, error) {
	file, err := os.Open(r.path)
	if err != nil {
		fmt.Printf("Error reading on file %s: %v", r.path, err)
		return nil, err
	}

	// Close the file after we're done
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var questions []question.Question
	for _, line := range lines {
		questions = append(questions, question.NewQuestion(line[0], line[1]))
	}

	return questions, nil
}
