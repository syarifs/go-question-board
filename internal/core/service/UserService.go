package service

import (
	"errors"
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

func (srv UserService) ReadUser() (res *[]response.UserList, err error) {
	var user *[]models.User
	user, err  = srv.repo.ReadUser()
	res, _ = utils.TypeConverter[[]response.UserList](&user)

	if utils.IsEmpty(res) {
		err = errors.New("Data Not Found")
	}

	
	return
}

func (srv UserService) ReadUserByID(id int) (res *response.UserDetails, err error) {
	var users *models.User
	users, err  = srv.repo.ReadUserByID(id)
	res, err = utils.TypeConverter[response.UserDetails](&users)

	if utils.IsEmpty(res) {
		err = errors.New("Data Not Found")
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
