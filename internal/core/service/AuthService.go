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

func (srv AuthService) Login(login request.LoginRequest) (res *response.UserDetails, err error) {
	var user models.User
	var checkPassword bool
	user, err  = srv.repo.Login(login.Email)

	err = errors.CheckError(user, err)

	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, user.Password)
		if checkPassword {
			res, _ = utils.TypeConverter[response.UserDetails](&user)
		} else {
			err = errors.New(417, "Wrong Password")
		}
	}

	return
}

func (srv AuthService) Logout(token models.Token) (err error) {
	err = srv.repo.RevokeToken(token)
	return
}

func (srv AuthService) CreateToken(id int, level string) (t models.Token, err error) {
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
	err = srv.repo.UpdateToken(tkn, new_token)
	token = new_token
	return
}
