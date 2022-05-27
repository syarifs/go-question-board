package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/errors"
)


type EvaluateService struct {
	repo repository.EvaluateRepository
}

func NewEvaluateService(repo repository.EvaluateRepository) *EvaluateService {
	return &EvaluateService{repo: repo}
}

func (srv EvaluateService) GetQuest(subject_id int, user request.User) (res *response.EvaluateQuestDetails, err error) {
	var quest *models.Questionnaire
	var subject *models.Subject

	class := utils.GetTagByName("Class", user.Tag)

	subject, quest, err = srv.repo.GetEvaluateQuest(subject_id, class)

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
	
	err = errors.CheckError(res, err)

	return
}

func (srv EvaluateService) Evaluate(req request.Answer, teacher_id, subject_id int, class string) (err error) {
	var answer []models.UserAnswer

	for _, v := range req.Answer {
		ans, _ := utils.TypeConverter[models.UserAnswer](&v)
		ans.UserID = req.User.ID
		ans.EvaluateTeacher.SubjectID = uint(subject_id)
		ans.EvaluateTeacher.TeacherID = uint(teacher_id)
		ans.EvaluateTeacher.Class = class
		answer = append(answer, *ans)
	}

	err = srv.repo.Evaluate(answer)
	
	err = errors.CheckError(nil, err)

	return
}

func (srv EvaluateService) ViewEvaluateResponse(teacher_id, subject_id int, class string) (res response.QuestResponses, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.GetEvaluateResponse(teacher_id, subject_id, class)

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
