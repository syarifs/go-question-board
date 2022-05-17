package mocktesting

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGORMSQLMock() (gdb *gorm.DB, mock sqlmock.Sqlmock, err error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	gdb, err = gorm.Open(mysql.New(mysql.Config{
			Conn: db,
			SkipInitializeWithVersion: true,
		}),
	)

	return
}
