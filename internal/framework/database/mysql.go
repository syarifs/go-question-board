package database

import (
	"fmt"
	"go-question-board/internal/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DB_USERNAME, utils.DB_PASSWORD, utils.DB_HOST, utils.DB_DATABASE,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
