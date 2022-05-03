package repository

import (
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
)

type questionnaireRepository struct {
	db *gorm.DB
}

func NewQuestionnaireRepository(db *gorm.DB) *questionnaireRepository {
	return &questionnaireRepository{db: db}
}

func (repo questionnaireRepository)CreateQuest(quest m.Questionnaire) (question m.Questionnaire, err error) {
	return
}

func (repo questionnaireRepository) ListMyQuest(user_id int) (quests []m.Questionnaire, err error) {
	return
}

func (repo questionnaireRepository) AvailableQuest(tag_id []int) (quests []m.Questionnaire, err error) {
	return
}

func (repo questionnaireRepository) UpdateQuest(id int, quest m.Questionnaire) (quests m.Questionnaire, err error) {
	return
}

func (repo questionnaireRepository) DeleteQuest(id int) (err error) {
	return
}

func (repo questionnaireRepository) ViewQuestByID(id int) (quest m.Questionnaire, err error) {
	return
}
