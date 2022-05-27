package repository

import m "go-question-board/internal/core/entity/models"

type EvaluateRepository interface {
	Evaluate([]m.UserAnswer) error
	GetEvaluateQuest(int, string) (*m.Subject, *m.Questionnaire, error)
	GetEvaluateResponse(int, int, string) (*m.Questionnaire, error)
}
