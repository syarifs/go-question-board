package request

import "go-question-board/internal/core/models"

type Answer struct {
	Questionnaire models.Questionnaire	`json:"questionnaire"`
	Answer []models.UserAnswer	`json:"answer"`
	User models.User `json:"user"`
}
