package response

import m "go-question-board/internal/core/entity/models"

type (
	User struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
		Status int `json:"status"`
		Level string `json:"level"`
		Major string `json:"major"`
	}

	UserDetails struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Password string `json:"-"`
		Name string `json:"name"`
		Status int `json:"status"`
		Role string `json:"role"`
		Major string `json:"major"`

		Tags []m.Tag `json:"tags" gorm:"many2many:user_tags"`
		// Subject []Subject `json:"subject"`
		// TeacherSubject []SubjectTeacher `json:"teacher_subject"`
	}
)
