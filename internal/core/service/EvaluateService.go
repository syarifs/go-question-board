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

func (srv EvaluateService) GetQuest() (res *response.AvailableQuestDetails, err error) {
	var quest *models.Questionnaire
	quest, err = srv.repo.GetEvaluateQuest()

	if utils.IsEmpty(quest) {
		err = errors.New("Data Not Found")
	}


	if err == nil {
		res, _ = utils.TypeConverter[response.AvailableQuestDetails](&quest)
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
