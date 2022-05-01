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

func (srv UserService) CreateUser(user models.UserModel) (res response.UserListResponse, err error) {
	err  = srv.repo.CreateUser(user)
	if err == nil {
		res = response.UserListResponse{
			Name: user.Name,
			Email: user.Email,
			Status: user.Status,
			Level: user.Level,
		}
	}
	return
}

func (srv UserService) ReadUser() (res []response.UserListResponse, err error) {
	var user *[]models.UserModel
	user, err  = srv.repo.ReadUser()
	if err == nil {
		for _, um := range *user {
			res = append(res, response.UserListResponse{
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

func (srv UserService) ReadUserByID(id int) (res response.UserDetailsResponse, err error) {
	var user *models.UserModel
	user, err  = srv.repo.ReadUserByID(id)
	if err == nil {
		res = response.UserDetailsResponse{
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

func (srv UserService) UpdateUser(id int, user models.UserModel) (res response.UserDetailsResponse, err error) {
	err  = srv.repo.UpdateUser(id, user)
	if err == nil {
		res = response.UserDetailsResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			Level: user.Level,
			Major: user.Major,
			Tags: user.Tags,
			Subject: user.Subject,
			Status: user.Status,
		}
	}
	return
}

func (srv UserService) DeleteUser(id int) (err error) {
	err  = srv.repo.DeleteUser(id)
	return
}
