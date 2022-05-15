package service

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
)

type SubjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

func (srv SubjectService) CreateSubject(subject request.SubjectRequest) (err error) {
	var sub *models.Subject
	sub, err = utils.TypeConverter[models.Subject](&subject)
	err  = srv.repo.CreateSubject(*sub)
	return
}

func (srv SubjectService) ReadSubject() (res *[]response.Subject, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadSubject()
	res, _ = utils.TypeConverter[[]response.Subject](&subject)

	if utils.IsEmpty(res) {
		err = errors.New("Data Not Found")
	}

	return
}

func (srv SubjectService) ReadUserSubject(user_id int) (res *[]response.Subject, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadSubjectByUserID(user_id)
	res, _ = utils.TypeConverter[[]response.Subject](&subject)

	if res == nil {
		err = errors.New("Data Not Found")
	}

	return
}

func (srv SubjectService) UpdateSubject(id int, subject request.SubjectRequest) (err error) {
	subject.ID = id
	sub, _ := utils.TypeConverter[models.Subject](&subject)

	for i := range sub.Teacher {
		sub.Teacher[i].SubjectID = subject.ID
	}

	err  = srv.repo.UpdateSubject(*sub)
	return
}

func (srv SubjectService) DeleteSubject(id int) (err error) {
	err  = srv.repo.DeleteSubject(id)
	return
}
