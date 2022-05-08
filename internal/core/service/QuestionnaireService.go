package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
)

type QuestionnaireService struct {
	repo repository.QuestionnaireRepository
}

func NewQuestionnaireService(repo repository.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{repo: repo}
}

func (repo *QuestionnaireService) CreateQuestionnaire(quest models.Questionnaire) (res response.QuestResponse, err error) {
	err = repo.repo.CreateQuest(quest)
	if err == nil {
		res.ID = quest.ID
		res.Title = quest.Title
		res.Description = quest.Description
		res.Tag = quest.Tags
		res.Question = quest.Question
	}
	return
}

func (repo *QuestionnaireService) MyQuestionnaire(user_id int) (res []response.MyQuestDahsboardResponse, err error) {
	var quest *[]models.Questionnaire
	quest, err = repo.repo.ListMyQuest(user_id)
	if err == nil {
		for _, v := range *quest {
			res = append(res, response.MyQuestDahsboardResponse{
				ID: v.ID,
				Title: v.Title,
				Description: v.Description,
				Tags: v.Tags,
				CreatedBy: int(v.CreatedBy),
				CountAnswered: len(v.Completor),
			})
		}
	}
	return
}

func (repo *QuestionnaireService) AvailableQuest(tags_id []models.Tag) (res []response.AvailableQuestionnareResponse, err error) {
	var tag_id []uint
	for _, t := range tags_id {
		tag_id = append(tag_id, t.ID)
	}
	var quest *[]models.Questionnaire
	quest, err = repo.repo.AvailableQuest(tag_id)
	if err == nil {
		for _, v := range *quest {
			res = append(res, response.AvailableQuestionnareResponse{
				ID: v.ID,
				Title: v.Title,
				Description: v.Description,
				Tags: v.Tags,
				CreatedBy: v.Creator,
			})
		}
	}

	return
}

func (repo *QuestionnaireService) UpdateQuest(id int, quest models.Questionnaire) (res response.QuestResponse, err error) {
	quest.ID = uint(id)
	err = repo.repo.UpdateQuest(quest)
	if err == nil {
		res.ID = quest.ID
		res.Title = quest.Title
		res.Description = quest.Description
		res.Tag = quest.Tags
		res.Question = quest.Question
	}
	return
}

func (repo *QuestionnaireService) DeleteQuest(id int) (err error) {
	err = repo.repo.DeleteQuest(id)
	return
}

func (repo *QuestionnaireService) ViewQuestByID(id int) (res response.QuestResponse, err error) {
	var quest *models.Questionnaire
	quest, err = repo.repo.ViewQuestByID(id)
	if err == nil {
		res.ID = quest.ID
		res.Title = quest.Title
		res.Description = quest.Description
		res.Tag = quest.Tags
		res.Question = quest.Question
	}
	return
}
