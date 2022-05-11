package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"reflect"
)

type QuestionnaireService struct {
	repo repository.QuestionnaireRepository
}

func NewQuestionnaireService(repo repository.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{repo: repo}
}

func (repo QuestionnaireService) CreateQuest(quest models.Questionnaire) (err error) {
	err = repo.repo.CreateQuest(quest)
	return
}

func (repo QuestionnaireService) MyQuest(user_id int) (res []response.QuestList, err error) {
	var quest *[]models.Questionnaire
	quest, err = repo.repo.MyQuest(user_id)
	if err == nil {
		for _, v := range *quest {
			res = append(res, response.QuestList{
				ID: v.ID,
				Title: v.Title,
				Description: v.Description,
				Tags: v.Tags,
			})
		}
	}
	return
}

func (repo QuestionnaireService) QuestForMe(tags []models.Tag) (res []response.AvailableQuestList, err error) {
	var tag_id []int
	var quest *[]models.Questionnaire

	for _, v := range tags {
		tag_id = append(tag_id, int(v.ID))
	}

	quest, err = repo.repo.QuestForMe(tag_id)
	if err == nil {
		for _, v := range *quest {
			if reflect.DeepEqual(v.Tags, tags) {
				res = append(res, response.AvailableQuestList{
					ID: v.ID,
					Title: v.Title,
					Description: v.Description,
					CreatedBy: v.Creator,
				})
			}
		}
	}

	return
}

func (repo QuestionnaireService) UpdateQuest(id int, quest models.Questionnaire) (err error) {
	quest.ID = uint(id)
	err = repo.repo.UpdateQuest(quest)
	return
}

func (repo QuestionnaireService) DeleteQuest(id int) (err error) {
	err = repo.repo.DeleteQuest(id)
	return
}

func (repo QuestionnaireService) ViewQuestResponse(id int) (res response.QuestResponses, err error) {
	var quest *models.Questionnaire
	quest, err = repo.repo.ViewQuestResponse(id)
	if err == nil {
		for _, v := range quest.Question {
			respondent := response.Respondent{
				QuestionID: v.ID,
				Question: v.Question,
				Response: v.UserResponse,
			}
			respondent.NumberRespondent = uint(len(v.UserResponse))
			res.Questions = append(res.Questions, respondent)
		}
		res.ID = quest.ID
		res.Title = quest.Title
	}
	return
}

func (repo QuestionnaireService) ViewQuestByID(id int) (res response.AvailabelQuestDetails, err error) {
	var quest *models.Questionnaire
	quest, err = repo.repo.ViewQuestByID(id)
	if err == nil {
		question := quest.Question
		for _, v := range question {
			res.Question = append(res.Question, response.Question{
				ID: v.ID,
				QuestionnaireID: v.QuestionnaireID,
				WithOption: v.WithOption,
				Question: v.Question,
				AnswerOption: v.AnswerOption,
			})
		}
		res.ID = quest.ID
		res.Title = quest.Title
		res.Description = quest.Description
		res.Tag = quest.Tags
	}
	return
}
