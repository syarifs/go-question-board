package response

import m "go-question-board/internal/core/models"

type (
	UserList struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
		Level m.Level
		Status int `json:"status"`
	}

	UserDetails struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		Name string `json:"name"`
		Tags []m.Tag `json:"tags"`
		Subject []interface{} `json:"subject"`
		Level m.Level `json:"level"`
		Major m.Major `json:"major"`
		Status int `json:"status"`
	}
)
