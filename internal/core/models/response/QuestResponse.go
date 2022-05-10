package response

import m "go-question-board/internal/core/models"

type QuestList struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []m.Tag `json:"tags" gorm:"many2many:questionnaire_tags"`
}

type QuestDetails struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []m.Tag `json:"tags" gorm:"many2many:questionnaire_tags"`
	CreatedBy m.User `json:"created_by"`
	Respondent m.User `json:"respondent"`
}

type AvailableQuestList struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedBy m.User `json:"created_by"`
}

type Question struct {
	ID uint `json:"id" gorm:"primaryKey"`
	QuestionnaireID uint `json:"questionnaire_id"`
	Question string `json:"question"`
	WithOption int `json:"with_option"`
	AnswerOption []m.AnswerOption `json:"answer_option"`
}

type AvailabelQuestDetails struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tag []m.Tag `json:"tags"`
	Question []Question `json:"questions"`
}

type Respondent struct {
	QuestionID uint `json:"question_id"`
	Question string `json:"question"`
	NumberRespondent uint `json:"number_of_response"`
	Response []m.UserAnswer `json:"answer"`
}

type QuestResponses struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Questions []Respondent `json:"questions"`
}
