package request

import m "go-question-board/internal/core/entity/models"

type User struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Level m.Level `json:"level"`
	Tag []m.Tag	`json:"tags"`
	Status int `json:"status"`
}

