package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/repository"
)

type SubjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

func (srv SubjectService) CreateSubject(subject models.Subject) (res models.Subject, err error) {
	err  = srv.repo.CreateSubject(subject)
	res = subject
	return
}

func (srv SubjectService) ReadSubject() (res []models.Subject, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadSubject()
	if err == nil {
		for _, um := range *subject {
			res = append(res, um)
		}
	}
	return
}

func (srv SubjectService) UpdateSubject(id int, subject models.Subject) (res models.Subject, err error) {
	subject.ID = uint(id)
	err  = srv.repo.UpdateSubject(subject)
	if err == nil {
		res = subject
	}
	return
}

func (srv SubjectService) DeleteSubject(id int) (err error) {
	err  = srv.repo.DeleteSubject(id)
	return
}
