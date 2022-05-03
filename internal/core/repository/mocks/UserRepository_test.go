package mocks

import (
	"errors"
	models "go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

var mockUser = &UserRepository{Mock: mock.Mock{}}
var userService = service.NewUserService(mockUser)

func TestCreateUser(t *testing.T) {}

func TestUpdateUser(t *testing.T) {}

func TestDeleteUser(t *testing.T) {}

func TestReadUser(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := []models.User{
			{
				Name: "Test",
				Email: "test@testify.io",
				Password: "test2",
				LevelID: 1,
			},
		}
		mockUser.On("ReadUser").Return(&data, nil).Once()
		user, err := userService.ReadUser()
		assert.NotNil(t, user)
		assert.NoError(t, err)
	})
	t.Run("Fail", func(t *testing.T) {
		mockUser.On("ReadUser").Return(&[]models.User{}, errors.New("fail")).Once()
		user, err := userService.ReadUser()
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

