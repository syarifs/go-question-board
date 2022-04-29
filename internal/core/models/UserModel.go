package models

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Status int `json:"status"`
	LevelID int `json:"level_id"`
	MajorID *int `json:"major_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Tags []TagModel `gorm:"many2many:user_tag"`
	Subject []SubjectModel `gorm:"many2many:user_subject"`
	Level LevelModel
	Major MajorModel
}

func (*UserModel) TableName() string {
	return "users"
}
