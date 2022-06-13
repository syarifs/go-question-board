package service

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/repository"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/errors"
	"go-question-board/internal/utils/jwt"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (srv AuthService) Login(login request.LoginRequest) (res response.User, err error) {
	var checkPassword bool
	res, err  = srv.repo.Login(login.Email)

	err = errors.CheckError(res, err)

	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, res.Password)
		if !checkPassword {
			err = errors.New(417, "Wrong Password")
		}
	}

	return
}

func (srv AuthService) Logout(token string) (err error) {
	err = srv.repo.RevokeToken(token)
	return
}

func (srv AuthService) CreateToken(id uint, level string) (t models.Token, err error) {
	t, err = jwt.CreateToken(float64(id), level)
	err = srv.repo.SaveToken(t)
	return
}

func (srv AuthService) RefreshToken(tkn models.Token) (token models.Token, err error) {

	if tkn.AccessToken == "" {
		err = errors.New(401, "Token Not Provided")
		return
	}

	new_token, err := jwt.RefreshToken(tkn)
	if err != nil {
		return
	}
	err = srv.repo.UpdateToken(tkn, new_token)
	token = new_token
	return
}
