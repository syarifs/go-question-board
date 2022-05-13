package repository

import (
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db: db}
}

func (repo tagRepository) CreateTag(tag m.Tag) (err error) {
	err = repo.db.Create(&tag).Error
	return
}

func (repo tagRepository) UpdateTag(tag m.Tag) (err error) {
	err = repo.db.Updates(&tag).Error
	return
}

func (repo tagRepository) DeleteTag(id int) (err error) {
	err = repo.db.Delete(&m.Tag{}, id).Error
	return
}

func (repo tagRepository) ReadTag() (tag *[]m.Tag, err error) {
	err = repo.db.Find(&tag).Error
	return
}
