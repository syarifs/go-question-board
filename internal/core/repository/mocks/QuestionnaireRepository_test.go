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

func TestUpdateQuestionnaire(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Questionnaire{
			ID: 1,
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
		mockQuestionnaire.On("UpdateQuest", data).Return(nil).Once()
		res, err := questionnaireService.UpdateQuest(1, data)
		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("UpdateQuest", mock.Anything).Return(errors.New("error")).Once()
		res, err := questionnaireService.UpdateQuest(0, models.Questionnaire{})
		assert.Empty(t, res)
		assert.Error(t, err)
	})
}

func TestDeleteQuestionnaire(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockQuestionnaire.On("DeleteQuest", 1).Return(nil).Once()
		err := questionnaireService.DeleteQuest(1)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("DeleteQuest", 1).Return(errors.New("fail")).Once()
		err := questionnaireService.DeleteQuest(1)
		assert.Error(t, err)
	})
}

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

func TestAvalaibleQuestionnaire(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := []models.Questionnaire{
			{
				Title: "Test Quest",
				Description: "Test Quest",
				Tags: []models.Tag{
					{ID: 1, Name: "Year", Value: "2019"},
				},
				Question: []models.Question{
					{Question: "Test Quest 1", WithOption: 0},
				},
				CreatedBy: 1,
			},
		}
		tag := []models.Tag{
			{ID: 1, Name: "Year", Value: "2019"},
		}
		mockQuestionnaire.On("AvailableQuest", tag).Return(&data, nil).Once()
		quest, err := questionnaireService.AvailableQuest(tag)
		assert.NotEmpty(t, quest)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		tag := []models.Tag{
			{ID: 1, Name: "Year", Value: "2019"},
		}
		mockQuestionnaire.On("AvailableQuest", tag).Return(nil, errors.New("fail")).Once()
		quest, err := questionnaireService.AvailableQuest(tag)
		assert.Empty(t, quest)
		assert.Error(t, err)
	})
}

func TestViewQuestionnaireByID(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Questionnaire{
			Title: "Test Quest",
			Description: "Test Quest",
			Tags: []models.Tag{
				{ID: 1, Name: "Year", Value: "2019"},
			},
			Question: []models.Question{
				{Question: "Test Quest 1", WithOption: 0},
			},
			CreatedBy: 1,
		}
		mockQuestionnaire.On("ViewQuestByID", 1).Return(&data, nil).Once()
		quest, err := questionnaireService.ViewQuestByID(1)
		assert.NotEmpty(t, quest)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("ViewQuestByID", 1).Return(&models.Questionnaire{}, errors.New("fail")).Once()
		quest, err := questionnaireService.ViewQuestByID(1)
		assert.Empty(t, quest)
		assert.Error(t, err)
	})
}
