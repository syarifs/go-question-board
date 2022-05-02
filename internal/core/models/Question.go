package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuestionnaireID uint `json:"questionnaire_id"`
	Question string `json:"question"`
	WithOption int `json:"with_option"`
	AnswerOption []AnswerOption `json:"answer_option"`
}

func (*Question) TableName() string {
	return "question"
}
