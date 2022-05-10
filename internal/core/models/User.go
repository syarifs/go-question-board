package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Status int `json:"status"`
	LevelID int `json:"level_id"`
	MajorID *int `json:"major_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Tags []Tag `json:"tags" gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subject []Subject `json:"subject" gorm:"many2many:user_subjects;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Level Level
	Major Major
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (*User) TableName() string {
	return "users"
}
