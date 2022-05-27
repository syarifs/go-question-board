package repository

import m "go-question-board/internal/core/entity/models"

type TagRepository interface {
	CreateTag(m.Tag) error
	UpdateTag(m.Tag) error
	DeleteTag(int) error
	ReadTag() (*[]m.Tag, error)
}
