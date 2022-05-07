package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/repository"
)

type QuestionnaireService struct {
	repo repository.QuestionnaireRepository
}

func NewQuestionnaireService(repo repository.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{repo: repo}
}

func (repo *QuestionnaireService) CreateQuestionnaire(quest models.Questionnaire) (res models.Questionnaire, err error) {
	err = repo.repo.CreateQuest(quest)
	if err == nil {
		res = quest
	}
	return
}

func (repo *QuestionnaireService) MyQuestionnaire(user_id int) (res *[]models.Questionnaire, err error) {
	res, err = repo.repo.ListMyQuest(user_id)
	return
}

func (repo *QuestionnaireService) AvailableQuest(tags_id []models.Tag) (res *[]models.Questionnaire, err error) {
	res, err = repo.repo.AvailableQuest(tags_id)
	return
}

func (repo *QuestionnaireService) UpdateQuest(id int, quest models.Questionnaire) (res models.Questionnaire, err error) {
	quest.ID = uint(id)
	err = repo.repo.UpdateQuest(quest)
	if err == nil {
		res = quest
	}
	return
}

func (repo *QuestionnaireService) DeleteQuest(id int) (err error) {
	err = repo.repo.DeleteQuest(id)
	return
}

func (repo *QuestionnaireService) ViewQuestByID(id int) (res *models.Questionnaire, err error) {
	res, err = repo.repo.ViewQuestByID(id)
	return
}
