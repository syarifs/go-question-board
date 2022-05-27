package mocks

import (
	"errors"
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/service"
	"go-question-board/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

var mockAuth = &AuthRepository{Mock: mock.Mock{}}
var authService = service.NewAuthService(mockAuth)

func TestLogin(t *testing.T) {
	password, _ := utils.HashPassword("test")
	t.Run("Succes", func(t *testing.T) {
		data := request.LoginRequest{
			Email: "test@testify.io",
			Password: "test",
		}

		ret := models.User{
			Email: data.Email,
			Password: password,
		}
		mockAuth.On("Login", data.Email).Return(ret, nil).Once()
		res, err := authService.Login(data)
		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := request.LoginRequest{}
		ret := models.User{}
		mockAuth.On("Login", "").Return(ret, errors.New("fail")).Once()
		res, err := authService.Login(data)
		assert.Empty(t, res)
		assert.Error(t, err)
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("Succes", func(t *testing.T) {
		data := models.Token{
			AccessToken: "lasgbfuilawfgbklabfgla",
			RefreshToken: "aklbtfgalwetbakldfadbff",
		}
		mockAuth.On("RefreshToken", data).Return(data, nil).Once()
		res, err := authService.RefreshToken(data)
		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := models.Token{}
		mockAuth.On("RefreshToken", data).Return(data, errors.New("fail")).Once()
		res, err := authService.RefreshToken(data)
		assert.Empty(t, res)
		assert.Error(t, err)
	})
}

