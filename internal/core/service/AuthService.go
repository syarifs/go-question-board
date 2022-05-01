package service

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/repository"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (srv AuthService) Login(login request.LoginRequest) (res response.UserDetailsResponse, err error) {
	var user models.UserModel
	user, err  = srv.repo.Login(login)
	res.ID = user.ID
	res.Email = user.Email
	res.Name = user.Name
	res.Level = user.Level
	res.Subject = user.Subject
	res.Tags = user.Tags
	res.Major = user.Major
	res.Status = user.Status
	return
}

func (srv AuthService) RefreshToken(str models.TokenModel) (token models.TokenModel, err error) {
	token, err = srv.repo.RefreshToken(str)
	return
}
