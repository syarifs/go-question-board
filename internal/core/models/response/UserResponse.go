package response

import m "go-question-board/internal/core/models"

type UserResponse struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Tags []m.TagModel `gorm:"many2many:user_tag"`
	Subject []m.SubjectModel `gorm:"many2many:user_subject"`
	Level m.LevelModel
	Major m.MajorModel
	Status int `json:"status"`
}
