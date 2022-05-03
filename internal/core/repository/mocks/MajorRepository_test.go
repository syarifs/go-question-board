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

func TestCreateMajor(t *testing.T) {}

func TestUpdateMajor(t *testing.T) {}

func TestDeleteMajor(t *testing.T) {}

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

