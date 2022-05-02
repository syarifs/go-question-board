package response

import m "go-question-board/internal/core/models"

type UserListResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Level m.Level
	Status int `json:"status"`
}

type UserDetailsResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Tags []m.Tag `json:"tags" gorm:"many2many:user_tag"`
	Subject []m.Subject `gorm:"many2many:user_subject"`
	Level m.Level
	Major m.Major
	Status int `json:"status"`
}
