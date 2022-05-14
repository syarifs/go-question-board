package database

import (
	"go-question-board/internal/core/models"

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
		models.UserAnswer{},
		models.TeacherSubject{},
		models.EvaluateTeacher{},
	)
	return
}
