package mocks

import (
	"errors"
	models "go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
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
		err := questionnaireService.CreateQuest(data)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("CreateQuest", mock.Anything).Return(errors.New("error")).Once()
		err := questionnaireService.CreateQuest(models.Questionnaire{})
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
		err := questionnaireService.UpdateQuest(1, data)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("UpdateQuest", mock.Anything).Return(errors.New("error")).Once()
		err := questionnaireService.UpdateQuest(0, models.Questionnaire{})
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
		mockQuestionnaire.On("MyQuest", 1).Return(&data, nil).Once()
		quest, err := questionnaireService.MyQuest(1)
		assert.NotEmpty(t, quest)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := []models.Questionnaire{}
		mockQuestionnaire.On("MyQuest", 1).Return(&data, errors.New("fail")).Once()
		quest, err := questionnaireService.MyQuest(1)
		assert.Empty(t, quest)
		assert.Error(t, err)
	})

	t.Run("Fail 0 integer", func(t *testing.T) {
		data := []models.Questionnaire{}
		mockQuestionnaire.On("MyQuest", nil).Return(&data, errors.New("fail")).Once()
		quest, err := questionnaireService.MyQuest(0)
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
		mockQuestionnaire.On("QuestForMe", 1,  []int{1}).Return(&data, nil).Once()
		quest, err := questionnaireService.QuestForMe(1, tag)
		assert.NotEmpty(t, quest)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		tag := []models.Tag{}
		mockQuestionnaire.On("QuestForMe", 1, []int(nil)).Return(nil, errors.New("fail")).Once()
		quest, err := questionnaireService.QuestForMe(1, tag)
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


func TestViewQuestResponse(t *testing.T) {
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
		mockQuestionnaire.On("ViewQuestResponse", 1).Return(&data, nil).Once()
		quest, err := questionnaireService.ViewQuestResponse(1)
		assert.NotEmpty(t, quest)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockQuestionnaire.On("ViewQuestResponse", 1).Return(&models.Questionnaire{}, errors.New("fail")).Once()
		quest, err := questionnaireService.ViewQuestResponse(1)
		assert.Empty(t, quest)
		assert.Error(t, err)
	})
}

func TestQuestAnswer(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		quest := models.Questionnaire{
			Title: "Test Quest",
			Description: "Test Quest",
			Tags: []models.Tag{
				{ID: 1, Name: "Year", Value: "2019"},
			},
			Question: []models.Question{
				{Question: "Test Quest 1", WithOption: 0},
			},
			CreatedBy: 1,
			Completor: []models.User{{ID: 1},},
		}
		userAnswer := []models.UserAnswer{
			{
				QuestionID: 1,
				Answer: "A",
				UserID: 1,
			},
			{
				QuestionID: 2,
				Answer: "A",
				UserID: 1,
			},
		}
		answer := request.Answer{
			Questionnaire: quest,
			Answer: []request.UserAnswer{
				{
					QuestionID: 1,
					Answer: "A",
				},
				{
					QuestionID: 2,
					Answer: "A",
				},
			},
			User: models.User{ID: 1},
		}
		mockQuestionnaire.On("Answer", quest, userAnswer).Return(nil).Once()
		err := questionnaireService.AnswerQuest(answer)
		assert.NoError(t, err)
	})
}
