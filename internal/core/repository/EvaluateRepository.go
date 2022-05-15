package repository

import m "go-question-board/internal/core/models"

type EvaluateRepository interface {
	Evaluate(m.Questionnaire, []m.UserAnswer) error
	GetEvaluateQuest() (*m.Questionnaire, error)
	GetEvaluateResponse(int) (*[]m.Questionnaire, error)
}
