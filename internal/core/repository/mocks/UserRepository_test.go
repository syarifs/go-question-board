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

func TestCreateUser(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.User{
			Name: "Test",
			Email: "test@testify.io",
			Password: "test2",
			LevelID: 1,
		}
		mockUser.On("CreateUser", data).Return(nil).Once()
		user, err := userService.CreateUser(data)
		assert.NotEmpty(t, user)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("CreateUser", models.User{}).Return(errors.New("fail")).Once()
		user, err := userService.CreateUser(models.User{})
		assert.Empty(t, user)
		assert.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.User{
			ID: 1,
			Name: "Test",
			Email: "test@testify.io",
			Password: "test2",
			LevelID: 1,
		}
		mockUser.On("UpdateUser", 1, data).Return(nil).Once()
		user, err := userService.UpdateUser(1, data)
		assert.NotEmpty(t, user)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("UpdateUser", 1, models.User{}).Return(errors.New("fail")).Once()
		user, err := userService.UpdateUser(1, models.User{})
		assert.Empty(t, user)
		assert.Error(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockUser.On("DeleteUser", 1).Return(nil).Once()
		err := userService.DeleteUser(1)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("DeleteUser", 1).Return(errors.New("fail")).Once()
		err := userService.DeleteUser(1)
		assert.Error(t, err)
	})
}

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
		assert.NotEmpty(t, user)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("ReadUser").Return(&[]models.User{}, errors.New("fail")).Once()
		user, err := userService.ReadUser()
		assert.Empty(t, user)
		assert.Error(t, err)
	})
}

