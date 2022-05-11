package mocks

import (
	"errors"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockTag = &TagRepository{Mock: mock.Mock{}}
var tagService = service.NewTagService(mockTag)

func TestCreateTag(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Tag{
			Name: "Year",
			Value: "2019",
		}
		mockTag.On("CreateTag", data).Return(nil).Once()
		err := tagService.CreateTag(data)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockTag.On("CreateTag", mock.Anything).Return(errors.New("error")).Once()
		err := tagService.CreateTag(models.Tag{})
		assert.Error(t, err)
	})

}

func TestUpdateTag(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := models.Tag{
			ID: 1,
			Name: "Year",
			Value: "2019",
		}
		mockTag.On("UpdateTag", data).Return(nil).Once()
		err := tagService.UpdateTag(1, data)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockTag.On("UpdateTag", mock.Anything).Return(errors.New("error")).Once()
		err := tagService.UpdateTag(1, models.Tag{})
		assert.Error(t, err)
	})
}

func TestDeleteTag(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		mockTag.On("DeleteTag", 1).Return(nil).Once()
		err := tagService.DeleteTag(1)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockTag.On("DeleteTag", 1).Return(errors.New("error")).Once()
		err := tagService.DeleteTag(1)
		assert.Error(t, err)
	})
}

func TestReadTag(t *testing.T) {
	t.Run("Sucess", func(t *testing.T) {
		data := []models.Tag{
			{
				Name: "Year",
				Value: "2019",
			},
		}
		mockTag.On("ReadTag").Return(&data, nil).Once()
		tag, err := tagService.ReadTag()
		assert.NotNil(t, tag)
		assert.NoError(t, err)
	})

	t.Run("Fail", func(t *testing.T) {
		mockTag.On("ReadTag").Return(nil, errors.New("error")).Once()
		tag, err := tagService.ReadTag()
		assert.Nil(t, tag)
		assert.Error(t, err)
	})
}

