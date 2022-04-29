package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initSQLite() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("./test_db/data.db"), &gorm.Config{})
	return
}
