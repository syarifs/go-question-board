package database

import (
	"go-question-board/internal/framework/database/seeds"

	"gorm.io/gorm"
)

func InitDatabase() (db *gorm.DB) {
	var err error
	db, err = initMySQL()
	if err != nil {
		panic(err)
	}
	err = migrateDB(db)
	if err != nil {
		panic(err)
	}
	seeds.NewSeeders(db)

	return db
}
