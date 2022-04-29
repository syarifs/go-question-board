package models

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Status string `json:"status"`
	LevelID int `json:"level_id"`

	Level LevelModel
}

func (*UserModel) TableName() string {
	return "users"
}
