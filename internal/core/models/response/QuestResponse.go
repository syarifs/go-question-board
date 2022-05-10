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

type AvailabelQuestDetails struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tag []m.Tag `json:"tags"`
	Question []m.Question `json:"questions"`
}
