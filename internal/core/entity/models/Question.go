package models

type Question struct {
	QuestionnaireID uint `json:"questionnaire_id" gorm:"primaryKey;autoIncrement:false"`
	Question string `json:"question"`
	UserResponse []UserAnswer `json:"user_response" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AnswerOption []AnswerOption `json:"answer_option" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (*Question) TableName() string {
	return "question"
}
