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


func (repo evaluateRepository) GetEvaluateQuest(subject_id int, class string) (subject *m.Subject,  quest *m.Questionnaire, err error) {

	err = repo.db.Debug().Preload(clause.Associations).
		Preload("Teacher", "class = ?", class).
		Preload("Teacher.User").
		First(&subject, "id = ?", subject_id).Error

	err = repo.db.Preload(clause.Associations).
		Preload("Question.AnswerOption").
		First(&quest, "type = 'Evaluate'").Error
	return
}

func (repo evaluateRepository) GetEvaluateResponse(user_id, subject_id int, class string) (quest *m.Questionnaire, err error) {

	var response_id []int

	repo.db.Model(m.EvaluateTeacher{}).
		Select("id").
		Where("teacher_id = ?", user_id).
		Where("subject_id = ?", subject_id).
		Where("class = ?", class).Scan(&response_id)

	err = repo.db.Debug().Preload(clause.Associations).
		Preload("Question.AnswerOption").
		Preload("Question.UserResponse", "evaluate_teacher_id IN ?", response_id).
		Preload("Question.UserResponse.User").
		Preload("Question.UserResponse.User.Level").
		First(&quest, "type = 'Evaluate'").Error
	return
}

func (repo evaluateRepository) Evaluate(ans []m.UserAnswer) (err error) {
	err = repo.db.Create(&ans).Error
	if err == nil {
		for _, et := range ans {
			err = repo.db.Create(&et.EvaluateTeacher).Error
			err = repo.db.Model(&ans).Association("EvaluateTeacher").Append(&et.EvaluateTeacher)
		}
	}
	return
}

