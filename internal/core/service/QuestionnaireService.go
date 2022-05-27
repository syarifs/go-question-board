package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/errors"
)

type QuestionnaireService struct {
	repo repository.QuestionnaireRepository
}

func NewQuestionnaireService(srv repository.QuestionnaireRepository) *QuestionnaireService {
	return &QuestionnaireService{repo: srv}
}

func (srv QuestionnaireService) CreateQuest(quest models.Questionnaire) (err error) {
	err = srv.repo.CreateQuest(quest)

	err = errors.CheckError(nil, err)

	return
}

func (srv QuestionnaireService) MyQuest(user_id int) (res *[]response.Quest, err error) {
	var quest *[]models.Questionnaire

	if err = errors.CheckError(user_id, nil); err != nil {
		return
	}
	
	quest, err = srv.repo.MyQuest(user_id)
	res, err = utils.TypeConverter[[]response.Quest](quest)

	err = errors.CheckError(res, err)

	return
}

func (srv QuestionnaireService) QuestForMe(id int, tags []models.Tag) (res []response.AvailableQuest, err error) {
	var tag_id []int
	var quest *[]models.Questionnaire

	if err = errors.CheckError(id, nil); err != nil {
		return
	}
	
	for _, v := range tags {
		tag_id = append(tag_id, int(v.ID))
	}

	quest, err = srv.repo.QuestForMe(id, tag_id)

	err = errors.CheckError(quest, err)

	for _, v := range *quest {
		if utils.TagEqual(tags, v.Tags) {
			que, _ := utils.TypeConverter[response.AvailableQuest](&v)
			res = append(res, *que)
		}
	}

	err = errors.CheckError(res, err)

	return
}

func (srv QuestionnaireService) UpdateQuest(id int, quest models.Questionnaire) (err error) {
	quest.ID = uint(id)
	err = srv.repo.UpdateQuest(quest)

	err = errors.CheckError(nil, err)

	return
}

func (srv QuestionnaireService) DeleteQuest(id int) (err error) {
	err = srv.repo.DeleteQuest(id)

	err = errors.CheckError(nil, err)

	return
}

func (srv QuestionnaireService) ViewQuestResponse(id int) (res response.QuestResponses, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.ViewQuestResponse(id)

	if err == nil {
		res.ID = quest.ID
		res.Title = quest.Title
		for _, v := range quest.Question {
			respondent, _ := utils.TypeConverter[response.Respondent](&v)
			respondent.NumberRespondent = uint(len(v.UserResponse))
			res.Questions = append(res.Questions, *respondent)
		}
	}

	err = errors.CheckError(res, err)

	return
}

func (srv QuestionnaireService) ViewQuestByID(id int) (res *response.AvailableQuestDetails, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.ViewQuestByID(id)
	
	err = errors.CheckError(quest, err)

	if err == nil {
		res, _ = utils.TypeConverter[response.AvailableQuestDetails](&quest)
	}

	return
}

func (srv QuestionnaireService) AnswerQuest(req request.Answer) (err error) {
	var answer []models.UserAnswer
	for _, v := range req.Answer {
		ans, _ := utils.TypeConverter[models.UserAnswer](&v)
		ans.UserID = req.User.ID
		answer = append(answer, *ans)
	}
	req.Questionnaire.Completor = []models.User{{ID: req.User.ID}}
	err = srv.repo.Answer(req.Questionnaire, answer)
	
	err = errors.CheckError(nil, err)

	return
}
