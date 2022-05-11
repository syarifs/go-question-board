package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
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
	user, err  = srv.repo.ReadUserByID(id)
	if err == nil {
		res = response.UserDetails{
			ID:      user.ID,
			Email:   user.Email,
			Name:    user.Name,
			Level:   user.Level,
			Status:  user.Status,
			Major: user.Major,
			Subject: user.Subject,
			Tags: user.Tags,
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
