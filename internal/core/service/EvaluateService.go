package service

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
)


type EvaluateService struct {
	repo repository.EvaluateRepository
}

func NewEvaluateService(repo repository.EvaluateRepository) *EvaluateService {
	return &EvaluateService{repo: repo}
}

func (srv EvaluateService) GetQuest(subject_id int, class string) (res *response.EvaluateQuestDetails, err error) {
	var quest *models.Questionnaire
	var subject *models.Subject
	subject, quest, err = srv.repo.GetEvaluateQuest(subject_id, class)

	if utils.IsEmpty(quest) {
		err = errors.New("Data Not Found")
	}


	if err == nil {
		res, _ = utils.TypeConverter[response.EvaluateQuestDetails](&quest)
		sub, _ := utils.TypeConverter[response.SubjectWithoutTeacher](&subject)
		res.Subject.ID = int(subject.ID)
		res.Subject = *sub
		for _, v := range subject.Teacher {
			res.Teacher.ID = int(v.User.ID)
			res.Teacher.Name = v.User.Name
		}
	}
	return
}

func (srv EvaluateService) Evaluate(req request.Answer, teacher_id, subject_id int) (err error) {
	var answer []models.UserAnswer

	for _, v := range req.Answer {
		ans, _ := utils.TypeConverter[models.UserAnswer](&v)
		ans.UserID = req.User.ID
		ans.EvaluateTeacher.SubjectID = uint(subject_id)
		ans.EvaluateTeacher.TeacherID = uint(teacher_id)
		answer = append(answer, *ans)
	}

	req.Questionnaire.Completor = []models.User{req.User}
	err = srv.repo.Evaluate(req.Questionnaire, answer)
	return
}

func (srv EvaluateService) ViewEvaluateResponse(teacher_id, subject_id int, class string) (res response.QuestResponses, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.GetEvaluateResponse(teacher_id, subject_id, class)
	
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
