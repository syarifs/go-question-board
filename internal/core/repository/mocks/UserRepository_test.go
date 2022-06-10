package mocks

import (
	"errors"
	"go-question-board/internal/core/entity/models"
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
			RoleID: 1,
		}
		mockUser.On("CreateUser", mock.Anything).Return(nil).Once()
		err := userService.CreateUser(data)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockUser.On("CreateUser", models.User{}).Return(errors.New("fail")).Once()
		err := userService.CreateUser(models.User{})
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
			RoleID: 1,
		}
		mockUser.On("UpdateUser", mock.Anything).Return(nil).Once()
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
				RoleID: 1,
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
			RoleID: 1,
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

