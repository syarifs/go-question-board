package response

import m "go-question-board/internal/core/models"

type MyQuestDahsboardResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []m.Tag `json:"tags" gorm:"many2many:questionnaire_tags"`
	CreatedBy int `json:"created_by"`
	CountAnswered int `json:"count_answered"`
}

type AvailableQuestionnareResponse struct {
	ID uint `json:"id"`
	QuestionnaireID uint `json:"questionnaire_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []m.Tag `json:"tags"`
	CreatedBy m.User `json:"created_by"`
}

type QuestResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tag []m.Tag `json:"tags"`
	Question []m.Question `json:"questions"`
}
