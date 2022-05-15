package response

import m "go-question-board/internal/core/models"

type (
	EvaluateRespondent struct {
		QuestionID uint `json:"question_id"`
		Question string `json:"question"`
		NumberRespondent uint `json:"number_of_response"`
		Response []m.UserAnswer `json:"answer"`
	}

	EvaluateQuestResponses struct {
		ID uint `json:"id"`
		Title string `json:"title"`
		Questions []EvaluateQuestResponses `json:"questions"`
	}
)
