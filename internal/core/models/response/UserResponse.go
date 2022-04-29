package response

import "go-question-board/internal/core/models"

type UserResponse struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Level models.LevelModel `json:"level"`
	Status string `json:"status"`
}
