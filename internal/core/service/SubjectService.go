package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/errors"
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

	err = errors.CheckError(nil, err)

	return
}

func (srv SubjectService) ReadSubject() (res *[]response.SubjectWithoutTeacher, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadSubject()
	res, _ = utils.TypeConverter[[]response.SubjectWithoutTeacher](&subject)

	err = errors.CheckError(res, err)

	return
}

func (srv SubjectService) ReadSubjectByID(id int) (res *response.Subject, err error) {
	var subject *models.Subject
	subject, err  = srv.repo.ReadSubjectByID(id)
	res, _ = utils.TypeConverter[response.Subject](&subject)
	
	err = errors.CheckError(res, err)

	return
}

func (srv SubjectService) ReadStudentSubject(id int) (res *[]response.SubjectWithTeacher, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadStudentSubject(id)
	res, _ = utils.TypeConverter[[]response.SubjectWithTeacher](&subject)

	err = errors.CheckError(res, err)

	return
}

func (srv SubjectService) ReadTeacherSubject(id int) (res *[]response.SubjectWithStudent, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadTeacherSubject(id)
	res, _ = utils.TypeConverter[[]response.SubjectWithStudent](&subject)

	err = errors.CheckError(res, err)

	return
}

func (srv SubjectService) UpdateSubject(id int, subject request.SubjectRequest) (err error) {
	subject.ID = uint(id)
	sub, _ := utils.TypeConverter[models.Subject](&subject)

	for i := range sub.Teacher {
		sub.Teacher[i].SubjectID = subject.ID
	}

	err  = srv.repo.UpdateSubject(*sub)
	err = errors.CheckError(nil, err)
	return
}

func (srv SubjectService) DeleteSubject(id int) (err error) {
	err  = srv.repo.DeleteSubject(id)
	err = errors.CheckError(nil, err)
	return
}
