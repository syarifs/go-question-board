package database

import (
	"go-question-board/internal/core/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		models.User{},
		models.Level{},
		models.Questionnaire{},
		models.Question{},
		models.AnswerOption{},
	)
	return
}
