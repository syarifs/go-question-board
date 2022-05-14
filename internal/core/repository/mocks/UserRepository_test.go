package mocks

import (
	"errors"
	models "go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"testing"

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
		userService.CreateUser(data)
		mockUser.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("CreateUser", models.User{}).Return(errors.New("fail")).Once()
		userService.CreateUser(models.User{})
		mockUser.AssertExpectations(t)
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
		mockUser.On("UpdateUser", data).Return(nil).Once()
		userService.UpdateUser(1, data)
		mockUser.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("UpdateUser", models.User{}).Return(errors.New("fail")).Once()
		userService.UpdateUser(0, models.User{})
		mockUser.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockUser.On("DeleteUser", 1).Return(nil).Once()
		userService.DeleteUser(1)
		mockUser.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("DeleteUser", 1).Return(errors.New("fail")).Once()
		userService.DeleteUser(1)
		mockUser.AssertExpectations(t)
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
		userService.ReadUser()
		mockUser.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("ReadUser").Return(&[]models.User{}, errors.New("fail")).Once()
		userService.ReadUser()
		mockUser.AssertExpectations(t)
	})
}

func TestReadUserByID(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.User{
			ID: 1,
			Name: "Test",
			Email: "test@testify.io",
			Password: "test2",
			LevelID: 1,
		}
		mockUser.On("ReadUserByID", 1).Return(&data, nil).Once()
		userService.ReadUserByID(1)
		mockUser.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("ReadUserByID", 1).Return(&models.User{}, errors.New("fail")).Once()
		userService.ReadUserByID(1)
		mockUser.AssertExpectations(t)
	})
}

