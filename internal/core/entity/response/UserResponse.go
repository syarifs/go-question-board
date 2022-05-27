package response

import m "go-question-board/internal/core/entity/models"

type (
	User struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
		Level m.Level `json:"level"`
		Status int `json:"status"`
	}

	UserDetails struct {
		ID        uint `gorm:"primary_key"`
		Email string `json:"email"`
		Name string `json:"name"`
		Status int `json:"status"`

		Tags []m.Tag `json:"tags"`
		Subject []Subject `json:"subject"`
		TeacherSubject []SubjectTeacher `json:"teacher_subject"`
		Level m.Level `json:"level"`
		Major m.Major `json:"major"`
	}
)
