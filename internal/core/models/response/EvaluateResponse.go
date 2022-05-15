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

	EvaluateQuestDetails struct {
		ID uint `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
		Teacher Teacher `json:"teacher"`
		Subject SubjectWithoutTeacher `json:"subject"`
		Tag []m.Tag `json:"tags"`
		Question []Question `json:"questions"`
	}
)
