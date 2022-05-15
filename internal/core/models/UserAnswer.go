package models

type UserAnswer struct {
	ID uint `json:"id" gorm:"primarykey"`
	Answer string `json:"answer"`
	QuestionID uint `json:"question_id"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
	EvaluateTeacherID *uint `json:"evaluate_teacher_id"`
	EvaluateTeacher EvaluateTeacher
}

func (*UserAnswer) TableName() string {
	return "user_answers"
}
