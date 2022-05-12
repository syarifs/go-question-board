package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
)

type SubjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

func (srv SubjectService) CreateSubject(subject request.SubjectRequest) (err error) {
	var sub models.Subject
	sub.Code = subject.Code
	sub.Name = subject.Name
	sub.Major = subject.Major
	for _, v := range subject.Teacher {
		sub.Teacher = append(sub.Teacher, models.TeacherSubject{
			UserID: v.ID,
			Class: v.Class,
		})
	}
	err  = srv.repo.CreateSubject(sub)
	return
}

func (srv SubjectService) ReadSubject() (res []response.Subject, err error) {
	var subject *[]models.Subject
	subject, err  = srv.repo.ReadSubject()
	if err == nil {
		for _, um := range *subject {
			sub :=  response.Subject{
				ID: int(um.ID),
				Code: um.Code,
				Name: um.Name,
				Major: um.Major,
			}

			for _, v := range um.Teacher {
				teacher := response.Teacher {
					ID: v.UserID,
					Name: v.User.Name,
					Class: v.Class,
				}
				sub.Teacher = append(sub.Teacher, teacher)
			}

			res = append(res, sub)
		}
	}
	return
}

func (srv SubjectService) UpdateSubject(id int, subject request.SubjectRequest) (err error) {
	var sub models.Subject
	sub.ID = uint(id)
	sub.Name = subject.Name
	sub.Code = subject.Code
	sub.Major = subject.Major
	for _, v := range subject.Teacher {
		sub.Teacher = append(sub.Teacher, models.TeacherSubject{
			SubjectID:  id,
			UserID: v.ID,
			Class: v.Class,
		})
	}
	err  = srv.repo.UpdateSubject(sub)
	return
}

func (srv SubjectService) DeleteSubject(id int) (err error) {
	err  = srv.repo.DeleteSubject(id)
	return
}
