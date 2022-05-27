package models

type AnswerOption struct {
	QuestionID uint `json:"question_id" gorm:"primaryKey;autoIncrement:false"`
	Answer string `json:"string_answer"`
}

func (*AnswerOption) TableName() string {
	return "answer_option"
}
