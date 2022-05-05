package repository

import (
	m "go-question-board/internal/core/models"
)

type QuestionnaireRepository interface {
	CreateQuest(m.Questionnaire) error
	ListMyQuest(int) (*[]m.Questionnaire, error)
	AvailableQuest([]int) (*[]m.Questionnaire, error)
	UpdateQuest(int, m.Questionnaire) error
	DeleteQuest(int) error
	ViewQuestByID(int) (*m.Questionnaire, error)
}
