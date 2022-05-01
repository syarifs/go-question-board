package response

import m "go-question-board/internal/core/models"

type UserListResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Level m.LevelModel
	Status int `json:"status"`
}

type UserDetailsResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Tags []m.TagModel `json:"tags" gorm:"many2many:user_tag"`
	Subject []m.SubjectModel `gorm:"many2many:user_subject"`
	Level m.LevelModel
	Major m.MajorModel
	Status int `json:"status"`
}
