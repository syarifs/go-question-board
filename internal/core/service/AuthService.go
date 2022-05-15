package service

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (srv AuthService) Login(login request.LoginRequest) (res *response.UserDetails, err error) {
	var user models.User
	var checkPassword bool
	user, err  = srv.repo.Login(login)

	if utils.IsEmpty(user) {
		err = errors.New("Data Not Found")
	}

	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, user.Password)
	}

	if checkPassword {
		res, _ = utils.TypeConverter[response.UserDetails](&user)
	} else {
		err = errors.New("Wrong Password")
	}

	return
}

func (srv AuthService) RefreshToken(str models.Token) (token models.Token, err error) {
	token, err = srv.repo.RefreshToken(str)
	return
}
