package mocks

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockMajor = &MajorRepository{Mock: mock.Mock{}}
var majorService = service.NewMajorService(mockMajor)

func TestCreateMajor(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Major{
			Code: "INF",
			Name: "Informatics",
		}
		mockMajor.On("CreateMajor", data).Return(nil).Once()
		major, err := majorService.CreateMajor(data)
		assert.NotEmpty(t, major)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockMajor.On("CreateMajor", mock.Anything).Return(errors.New("error")).Once()
		major, err := majorService.CreateMajor(models.Major{})
		assert.Empty(t, major)
		assert.Error(t, err)
	})

}

func TestUpdateMajor(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Major{
			ID: 1,
			Code: "INF",
			Name: "Informatics",
		}
		mockMajor.On("UpdateMajor", data).Return(nil).Once()
		major, err := majorService.UpdateMajor(1, data)
		assert.NotEmpty(t, major)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockMajor.On("UpdateMajor", mock.Anything).Return(errors.New("error")).Once()
		major, err := majorService.UpdateMajor(1, models.Major{})
		assert.Empty(t, major)
		assert.Error(t, err)
	})
}

func TestDeleteMajor(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockMajor.On("DeleteMajor", 1).Return(nil).Once()
		err := majorService.DeleteMajor(1)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockMajor.On("DeleteMajor", 1).Return(errors.New("error")).Once()
		err := majorService.DeleteMajor(1)
		assert.Error(t, err)
	})
}

func TestReadMajor(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := []models.Major{
			{Code: "INF", Name: "Informatics"},
		}
		mockMajor.On("ReadMajor").Return(&data, nil).Once()
		major, err := majorService.ReadMajor()
		assert.NotNil(t, major)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockMajor.On("ReadMajor").Return(nil, errors.New("error")).Once()
		major, err := majorService.ReadMajor()
		assert.Nil(t, major)
		assert.Error(t, err)
	})
}

