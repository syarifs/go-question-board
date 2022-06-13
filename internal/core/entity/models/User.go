package models

import "time"

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Email string `json:"email" validator:"required"`
	Password string `json:"password" validator:"required"`
	Status int `json:"status"`
	RoleID int `json:"role_id" validator:"required"`
	Role Role

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
