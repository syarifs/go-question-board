package models

// import "github.com/jinzhu/gorm"


type Role struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*Role) TableName() string {
	return "roles"
}
