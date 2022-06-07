package database

import (
	"fmt"
	"go-question-board/internal/utils/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
