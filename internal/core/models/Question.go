package models

type Question struct {
	ID uint `json:"id" gorm:"primaryKey"`
	QuestionnaireID uint `json:"questionnaire_id"`
	Question string `json:"question"`
	WithOption int `json:"with_option"`
	UserResponse []UserAnswer `json:"user_response" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnswerOption []AnswerOption `json:"answer_option" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (*Question) TableName() string {
	return "question"
}
