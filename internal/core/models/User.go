package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Status int `json:"status"`
	LevelID int `json:"level_id"`
	MajorID *int `json:"major_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	Tags []Tag `json:"tags" gorm:"many2many:user_tag"`
	Subject []Subject `json:"subject" gorm:"many2many:user_subject"`
	Level Level
	Major Major
}

func (*User) TableName() string {
	return "users"
}
