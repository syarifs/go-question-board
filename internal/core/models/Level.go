package models

// import "github.com/jinzhu/gorm"


type Level struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*Level) TableName() string {
	return "roles"
}
