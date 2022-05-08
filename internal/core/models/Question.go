package models

type Question struct {
	ID uint `json:"id" gorm:"primaryKey"`
	QuestionnaireID uint `json:"questionnaire_id"`
	Question string `json:"question"`
	WithOption int `json:"with_option"`
	AnswerOption []AnswerOption `json:"answer_option"`
}

func (*Question) TableName() string {
	return "question"
}
