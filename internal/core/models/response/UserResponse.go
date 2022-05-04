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
	Tags []m.Tag `json:"tags"`
	Subject []m.Subject `json:"subject"`
	Level m.Level `json:"level"`
	Major m.Major `json:"major"`
	Status int `json:"status"`
}
