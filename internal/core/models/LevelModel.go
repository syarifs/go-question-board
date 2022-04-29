package models

// import "github.com/jinzhu/gorm"


type LevelModel struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*LevelModel) TableName() string {
	return "roles"
}
