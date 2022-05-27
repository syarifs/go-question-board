package database

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		models.Major{},
		models.Subject{},
		models.Level{},
		models.User{},
		models.Questionnaire{},
		models.Question{},
		models.AnswerOption{},
		models.TeacherSubject{},
		models.EvaluateTeacher{},
		models.UserAnswer{},
	)
	return
}
