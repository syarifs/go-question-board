package database

import (
	"go-question-board/internal/core/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		models.UserModel{},
		models.LevelModel{},
	)
	return
}
