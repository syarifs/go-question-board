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

func (srv AuthService) Login(login request.LoginRequest) (res response.UserDetails, err error) {
	var user models.User
	var checkPassword bool
	user, err  = srv.repo.Login(login)
	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, user.Password)
	}
	if checkPassword {
		res.ID = user.ID
		res.Email = user.Email
		res.Name = user.Name
		res.Level = user.Level
		res.Subject = user.Subject
		res.Tags = user.Tags
		res.Major = user.Major
		res.Status = user.Status
	} else {
		err = errors.New("Wrong Password")
	}
	return
}

func (srv AuthService) RefreshToken(str models.Token) (token models.Token, err error) {
	token, err = srv.repo.RefreshToken(str)
	return
}
