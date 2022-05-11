package service

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"reflect"
)

type QuestionnaireService struct {
	repo repository.QuestionnaireRepository
}

func NewQuestionnaireService(srv repository.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{repo: srv}
}

func (srv QuestionnaireService) CreateQuest(quest models.Questionnaire) (err error) {
	err = srv.repo.CreateQuest(quest)
	return
}

func (srv QuestionnaireService) MyQuest(user_id int) (res []response.QuestList, err error) {
	var quest *[]models.Questionnaire

	if user_id == 0 {
		err = errors.New("User ID not provided")
		return
	}

	quest, err = srv.repo.MyQuest(user_id)
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

func (srv QuestionnaireService) QuestForMe(tags []models.Tag) (res []response.AvailableQuestList, err error) {
	var tag_id []int
	var quest *[]models.Questionnaire

	for _, v := range tags {
		tag_id = append(tag_id, int(v.ID))
	}

	quest, err = srv.repo.QuestForMe(tag_id)
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

func (srv QuestionnaireService) UpdateQuest(id int, quest models.Questionnaire) (err error) {
	quest.ID = uint(id)
	err = srv.repo.UpdateQuest(quest)
	return
}

func (srv QuestionnaireService) DeleteQuest(id int) (err error) {
	err = srv.repo.DeleteQuest(id)
	return
}

func (srv QuestionnaireService) ViewQuestResponse(id int) (res response.QuestResponses, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.ViewQuestResponse(id)
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

func (srv QuestionnaireService) ViewQuestByID(id int) (res response.AvailabelQuestDetails, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.ViewQuestByID(id)
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

func (srv QuestionnaireService) AnswerQuest(req request.Answer) (err error) {
	var answer []models.UserAnswer
	for _, v := range req.Answer {
		ans := models.UserAnswer{
			ID: v.ID,
			QuestionID: v.QuestionID,
			Answer: v.Answer,
			UserID: req.User.ID,
		}
		answer = append(answer, ans)
	}
	req.Questionnaire.Completor = []models.User{req.User}
	err = srv.repo.Answer(req.Questionnaire, answer)
	return
}
