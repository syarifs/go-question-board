package repository

import (
	"go-question-board/internal/core/repository"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *repository.Repository {
	return &repository.Repository{
		Auth: NewAuthRepository(db),
		User: NewUserRepository(db),
		Tag: NewTagRepository(db),
		Subject: NewSubjectRepository(db),
		Major: NewMajorRepository(db),
		Questionnaire: NewQuestionnaireRepository(db),
		EvaluateTeacher: NewEvaluateRepository(db),
	}
}
