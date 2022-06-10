package database

import (
	"go-question-board/internal/framework/database/seeds"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func InitDatabase() (sqldb *gorm.DB, mongodb *mongo.Database) {
	var err error

	mongodb = initMongoDB()

	sqldb, err = initMySQL()
	if err != nil {
		panic(err)
	}

	err = migrateDB(sqldb)
	if err != nil {
		panic(err)
	}

	seeds.NewSeeders(sqldb)

	return
}
