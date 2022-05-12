package service

import (
	"fmt"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (srv UserService) CreateUser(user models.User) (err error) {
	err  = srv.repo.CreateUser(user)
	return
}

func (srv UserService) ReadUser() (res []response.UserList, err error) {
	var user *[]models.User
	user, err  = srv.repo.ReadUser()
	if err == nil {
		for _, um := range *user {
			res = append(res, response.UserList{
				ID: um.ID,
				Name: um.Name,
				Email: um.Email,
				Level: um.Level,
				Status: um.Status,
			})
		}
	}
	return
}

func (srv UserService) ReadUserByID(id int) (res response.UserDetails, err error) {
	var user *models.User
	var subject response.UserSubject

	user, err  = srv.repo.ReadUserByID(id)
	class := utils.GetTagByName("Class", user.Tags)

	if err == nil {
		res = response.UserDetails{
			ID:      user.ID,
			Email:   user.Email,
			Name:    user.Name,
			Level:   user.Level,
			Status:  user.Status,
			Major: user.Major,
			Tags: user.Tags,
		}

		if user.Level.Name == "Teacher" {
			for _, v := range user.TeacherSubject {
				subject.ID = v.SubjectID
				subject.Code = v.Subject.Code
				subject.Name = v.Subject.Name
				subject.Major = v.Subject.Major

				subject.Teacher.ID = v.UserID
				subject.Teacher.Name = v.User.Name
				subject.Teacher.Class = v.Class

				res.Subject = append(res.Subject, subject)
			}
		} else {
			for _, v := range user.Subject {
				subject.ID = int(v.ID)
				subject.Code = v.Code
				subject.Name = v.Name
				subject.Major = v.Major

				for _, val := range v.Teacher {
					if val.Class == class {
						subject.Teacher.ID = val.UserID
						subject.Teacher.Name = val.User.Name
						subject.Teacher.Class = val.Class
					}
				}

				fmt.Println(subject)
				res.Subject = append(res.Subject, subject)
			}
		}
	}
	return
}

func (srv UserService) UpdateUser(id int, user models.User) (err error) {
	user.ID = uint(id)
	err  = srv.repo.UpdateUser(user)
	return
}

func (srv UserService) DeleteUser(id int) (err error) {
	err  = srv.repo.DeleteUser(id)
	return
}
