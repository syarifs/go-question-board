package repository

import (
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type evaluateRepository struct {
	db *gorm.DB
}

func NewEvaluateRepository(db *gorm.DB) *evaluateRepository {
	return &evaluateRepository{db: db}
}


func (repo evaluateRepository) GetEvaluateQuest() (quest *m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).
		Preload("Question.AnswerOption").
		Find(&quest, "type = 'Evaluate'").Error
	return
}

func (repo evaluateRepository) GetEvaluateResponse(user_id int) (quest *[]m.Questionnaire, err error) {
	err = repo.db.Preload(clause.Associations).Find(&quest, "type = 'Evaluate'").Error
	return
}

func (repo evaluateRepository) Evaluate(quest m.Questionnaire, ans []m.UserAnswer) (err error) {
	err = repo.db.Create(&ans).Error
	if err == nil {
		err = repo.db.Model(&quest).Association("Completor").Append(&quest.Completor)
	}
	return
}

