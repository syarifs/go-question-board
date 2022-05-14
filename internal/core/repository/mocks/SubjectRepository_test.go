package mocks

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/service"
	"go-question-board/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockSubject = &SubjectRepository{Mock: mock.Mock{}}
var subjectService = service.NewSubjectService(mockSubject)

func TestCreateSubject(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Subject{
			Code: "INF",
			Name: "Informatics",
		}
		mockSubject.On("CreateSubject", data).Return(nil).Once()
		req, _ := utils.TypeConverter[request.SubjectRequest](&data)
		err := subjectService.CreateSubject(*req)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockSubject.On("CreateSubject", mock.Anything).Return(errors.New("error")).Once()
		err := subjectService.CreateSubject(request.SubjectRequest{})
		assert.Error(t, err)
	})

}

func TestUpdateSubject(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Subject{
			ID: 1,
			Code: "INF",
			Name: "Informatics",
		}
		mockSubject.On("UpdateSubject", data).Return(nil).Once()
		req, _ := utils.TypeConverter[request.SubjectRequest](&data)
		err := subjectService.UpdateSubject(1, *req)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockSubject.On("UpdateSubject", mock.Anything).Return(errors.New("error")).Once()
		err := subjectService.UpdateSubject(1, request.SubjectRequest{})
		assert.Error(t, err)
	})
}

func TestDeleteSubject(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockSubject.On("DeleteSubject", 1).Return(nil).Once()
		err := subjectService.DeleteSubject(1)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockSubject.On("DeleteSubject", 1).Return(errors.New("error")).Once()
		err := subjectService.DeleteSubject(1)
		assert.Error(t, err)
	})
}

func TestReadSubject(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := []models.Subject{
			{Code: "INF", Name: "Informatics"},
		}
		mockSubject.On("ReadSubject").Return(&data, nil).Once()
		subject, err := subjectService.ReadSubject()
		assert.NotNil(t, subject)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockSubject.On("ReadSubject").Return(nil, errors.New("error")).Once()
		subject, err := subjectService.ReadSubject()
		assert.Nil(t, subject)
		assert.Error(t, err)
	})
}

