package database

import (
	"go-question-board/internal/framework/database/seeds"

	"gorm.io/gorm"
)

func InitDatabase(driver string) (db *gorm.DB) {
	var err error
	if driver == "mysql" {
		db, err = initMySQL()
	} else if driver == "sqlite" {
		db, err = initSQLite()
	}
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
