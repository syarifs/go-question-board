package mocks

import (
	"errors"
	models "go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

var mockQuestionnaire = &QuestionnaireRepository{Mock: mock.Mock{}}
var questionnaireService = service.NewQuestionnaireService(mockQuestionnaire)

func TestCreateQuestionnaire(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Questionnaire{
			Title: "Test Quest",
			Description: "Test Quest",
			Tags: []models.Tag{
				{Name: "Year", Value: "2019"},
			},
			Question: []models.Question{
				{Question: "Test Quest 1", WithOption: 0},
			},
			CreatedBy: 1,
		}
		mockQuestionnaire.On("CreateQuest", data).Return(nil).Once()
		questionnaire, err := questionnaireService.CreateQuestionnaire(data)
		assert.NotEmpty(t, questionnaire)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("CreateQuest", mock.Anything).Return(errors.New("error")).Once()
		questionnaire, err := questionnaireService.CreateQuestionnaire(models.Questionnaire{})
		assert.Empty(t, questionnaire)
		assert.Error(t, err)
	})
}

// func TestUpdateQuestionnaire(t *testing.T) {
// 	t.Run("Sucess", func(t *testing.T) {
// 		data := models.Questionnaire{
// 			ID: 1,
// 			Code: "INF",
// 			Name: "Informatics",
// 		}
// 		mockQuestionnaire.On("UpdateQuestionnaire", data).Return(nil).Once()
// 		questionnaire, err := questionnaireService.UpdateQuestionnaire(1, data)
// 		assert.NotEmpty(t, questionnaire)
// 		assert.NoError(t, err)
// 	})
//
// 	t.Run("Fail", func(t *testing.T) {
// 		mockQuestionnaire.On("UpdateQuestionnaire", mock.Anything).Return(errors.New("error")).Once()
// 		questionnaire, err := questionnaireService.UpdateQuestionnaire(1, models.Questionnaire{})
// 		assert.Empty(t, questionnaire)
// 		assert.Error(t, err)
// 	})
// }

// func TestDeleteQuestionnaire(t *testing.T) {}

func TestListMyQuestionnaire(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := []models.Questionnaire{
			{
				Title: "Test Quest",
				Description: "Test Quest",
				Tags: []models.Tag{
					{Name: "Year", Value: "2019"},
				},
				Question: []models.Question{
					{Question: "Test Quest 1", WithOption: 0},
				},
				CreatedBy: 1,
			},
		}
		mockQuestionnaire.On("ListMyQuest", 1).Return(&data, nil).Once()
		quest, err := questionnaireService.MyQuestionnaire(1)
		assert.NotEmpty(t, quest)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := []models.Questionnaire{}
		mockQuestionnaire.On("ListMyQuest", 1).Return(&data, errors.New("fail")).Once()
		quest, err := questionnaireService.MyQuestionnaire(1)
		assert.Empty(t, quest)
		assert.Error(t, err)
	})
}

// func TestAvalaibleQuestionnaire(t *testing.T) {}

// func TestViewQuestionnaireByID(t *testing.T) {}
