package mocks

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/service"
	"go-question-board/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

var mockEval = &EvaluateRepository{Mock: mock.Mock{}}
var evalService = service.NewEvaluateService(mockEval)

func TestEvaluate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data := []models.UserAnswer{
			{
				Answer: "A",
				QuestionID: 1,
				UserID: uint(1),
				EvaluateTeacher: models.EvaluateTeacher{
					SubjectID: 1,
					TeacherID: 1,
					Class: "A",
				},
			},
		}
		reqAns, _ := utils.TypeConverter[[]request.UserAnswer](&data)
		req := request.Answer{
			Questionnaire: models.Questionnaire{ID: 1},
			Answer: *reqAns,
			User: models.User{ID: 1},
			
		}
		mockEval.On("Evaluate", data).Return(nil).Once()
		err := evalService.Evaluate(req, 1, 1, "A")
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := []models.UserAnswer{
			{
				UserID: uint(1),
				EvaluateTeacher: models.EvaluateTeacher{
					SubjectID: 1,
					TeacherID: 1,
					Class: "A",
				},
			},
		}
		reqAns, _ := utils.TypeConverter[[]request.UserAnswer](&data)
		req := request.Answer{
			Questionnaire: models.Questionnaire{ID: 1},
			Answer: *reqAns,
			User: models.User{ID: 1},
			
		}
		mockEval.On("Evaluate", data).Return(errors.New("fail")).Once()
		err := evalService.Evaluate(req, 1, 1, "A")
		assert.Error(t, err)
	})
}

func TestGetEvaluateQuest(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		quest := models.Questionnaire{
			ID: 1,
			Question: []models.Question{
				{
					ID: 1,
					Question: "Test",
					QuestionnaireID: 1,
					AnswerOption: []models.AnswerOption{
						{
							ID: 1,
							Answer: "A",
							QuestionID: 1,
						},
					},
				},
			},
		}
		subject := models.Subject{
			ID: 1,
			Code: "TST",
			Name: "Test",
			Teacher: []models.TeacherSubject{
				{
					ID: 1,
					SubjectID: 1,
					User: models.User{ID: 1},
				},
			},
		}
		mockEval.On("GetEvaluateQuest", 1, "A").Return(&subject, &quest, nil).Once()
		sub, err := evalService.GetQuest(1, "A")
		assert.NotEmpty(t, sub)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		quest := models.Questionnaire{}
		subject := models.Subject{}
		mockEval.On("GetEvaluateQuest", 1, "A").Return(&subject, &quest, errors.New("fail")).Once()
		sub, err := evalService.GetQuest(1, "A")
		assert.Empty(t, sub)
		assert.Error(t, err)
	})
}

func TestGetEvaluateResponse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		quest := models.Questionnaire{
			ID: 1,
			Question: []models.Question{
				{
					ID: 1,
					Question: "Test",
					QuestionnaireID: 1,
					AnswerOption: []models.AnswerOption{
						{
							ID: 1,
							Answer: "A",
							QuestionID: 1,
						},
					},
				},
			},
		}
		mockEval.On("GetEvaluateResponse", 1, 1, "A").Return(&quest, nil).Once()
		sub, err := evalService.ViewEvaluateResponse(1, 1, "A")
		assert.NotEmpty(t, sub)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		quest := models.Questionnaire{}
		mockEval.On("GetEvaluateResponse", 1, 1, "A").Return(&quest, errors.New("fail")).Once()
		sub, err := evalService.ViewEvaluateResponse(1, 1, "A")
		assert.Empty(t, sub)
		assert.Error(t, err)
	})
}
