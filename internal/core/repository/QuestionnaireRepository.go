package repository

import (
	m "go-question-board/internal/core/models"
)

type QuestionnaireRepository interface {
	CreateQuest(m.Questionnaire) error
	MyQuest(int) (*[]m.Questionnaire, error)
	UpdateQuest(m.Questionnaire) error
	DeleteQuest(int) error
	ViewQuestByID(int) (*m.Questionnaire, error)
	ViewQuestResponse(int) (*m.Questionnaire, error)
	QuestForMe([]uint) (*[]m.Questionnaire, error)
	Answer(m.Questionnaire, []m.UserAnswer) error
}
