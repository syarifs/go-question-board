package models

type UserAnswer struct {
	ID uint `json:"id" gorm:"foreignKey"`
	AnswerOptionID *uint `json:"answer_option_id"`
	Answer string `json:"answer"`
	QuestionID uint `json:"question_id"`
	UserID uint `json:"user_id"`

	AnswerOption AnswerOption
	Question Question
	User User
}

func (*UserAnswer) TableName() string {
	return "user_answers"
}
