package response

import m "go-question-board/internal/core/models"

type QuestDahsboardResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []m.Tag `json:"tags" gorm:"many2many:questionnaire_tags"`
	CreatedBy int `json:"created_by"`
	AnsweredUser int `json:"answered_user"`
}

type QuestResponse struct {
	QuestionnaireID uint `json:"questionnaire_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Question []m.Question `json:"questions"`
	AnsweredUser int `json:"answered_user"`
}
