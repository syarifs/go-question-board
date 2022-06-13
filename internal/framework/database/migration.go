package database

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		models.User{},
		models.Profile{},
	)
	return
}
