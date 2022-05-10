package mocks

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/service"
	"go-question-board/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

var mockAuth = &AuthRepository{Mock: mock.Mock{}}
var authService = service.NewAuthService(mockAuth)

func TestLogin(t *testing.T) {
	t.Run("Succes", func(t *testing.T) {
		data := request.LoginRequest{
			Email: "test@testify.io",
			Password: "test",
		}
		ret := models.User{
			Email: data.Email,
			Password: utils.HashPassword(data.Password),
		}
		mockAuth.On("Login", data).Return(ret, nil).Once()
		auth, err := authService.Login(data)
		assert.NotEmpty(t, auth)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := request.LoginRequest{}
		ret := models.User{}
		mockAuth.On("Login", data).Return(ret, errors.New("fail")).Once()
		auth, err := authService.Login(data)
		assert.Empty(t, auth)
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
		auth, err := authService.RefreshToken(data)
		assert.NotEmpty(t, auth)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		data := models.Token{}
		mockAuth.On("RefreshToken", data).Return(data, errors.New("fail")).Once()
		auth, err := authService.RefreshToken(data)
		assert.Empty(t, auth)
		assert.Error(t, err)
	})
}

