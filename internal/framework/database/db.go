package database

import "gorm.io/gorm"

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

	return db
}
