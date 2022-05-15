package service

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
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

	if utils.IsEmpty(quest) {
		err = errors.New("Data Not Found")
	}


	if err == nil {
		for _, v := range *quest {
			que, _ := utils.TypeConverter[response.QuestList](&v)
			res = append(res, *que)
		}
	}
	return
}

func (srv QuestionnaireService) QuestForMe(id int, tags []models.Tag) (res []response.AvailableQuestList, err error) {
	var tag_id []int
	var quest *[]models.Questionnaire

	for _, v := range tags {
		tag_id = append(tag_id, int(v.ID))
	}

	quest, err = srv.repo.QuestForMe(id, tag_id)

	if utils.IsEmpty(quest) {
		err = errors.New("Data Not Found")
	}

	if err == nil {
		for _, v := range *quest {
			if reflect.DeepEqual(v.Tags, tags) {
				que, _ := utils.TypeConverter[response.AvailableQuestList](&v)
				res = append(res, *que)
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
	
	if utils.IsEmpty(quest) {
		err = errors.New("Data Not Found")
	}

	if err == nil {
		res.ID = quest.ID
		res.Title = quest.Title
		for _, v := range quest.Question {
			respondent, _ := utils.TypeConverter[response.Respondent](&v)
			respondent.NumberRespondent = uint(len(v.UserResponse))
			res.Questions = append(res.Questions, *respondent)
		}
	}
	return
}

func (srv QuestionnaireService) ViewQuestByID(id int) (res *response.AvailableQuestDetails, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.ViewQuestByID(id)
	
	if utils.IsEmpty(quest) {
		err = errors.New("Data Not Found")
	}

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
	req.Questionnaire.Completor = []models.User{req.User}
	err = srv.repo.Answer(req.Questionnaire, answer)
	return
}
