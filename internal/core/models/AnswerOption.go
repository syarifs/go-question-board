package models

type AnswerOption struct {
	ID uint `json:"id" gorm:"primarykey"`
	QuestionID uint `json:"question_id"`
	Answer string `json:"string"`
}

func (*AnswerOption) TableName() string {
	return "answer_option"
}
