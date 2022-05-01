package repository

import (
	m "go-question-board/internal/core/models"
)

type TagRepository interface {
	CreateTag(m.TagModel) error
	UpdateTag(int, m.TagModel) error
	DeleteTag(int) error
	ReadTag() (*[]m.TagModel, error)
}
