package request

import "go-question-board/internal/core/models"

type UserAnswer struct {
	ID uint `json:"id" gorm:"primarykey"`
	Answer string `json:"answer"`
	QuestionID uint `json:"question_id"`
}

type Answer struct {
	Questionnaire models.Questionnaire	`json:"quest"`
	Answer []UserAnswer	`json:"answers"`
	User models.User `json:"user"`
}
