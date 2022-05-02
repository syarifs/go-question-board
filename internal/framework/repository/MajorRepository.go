package repository

import (
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
)

type majorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) *majorRepository {
	return &majorRepository{db: db}
}

func (repo majorRepository) CreateMajor(major m.Major) (err error) {
	err = repo.db.Create(&major).Error
	return
}

func (repo majorRepository) UpdateMajor(id int, major m.Major) (err error) {
	err = repo.db.Where("id = ?", id).Updates(&major).Error
	return
}

func (repo majorRepository) DeleteMajor(id int) (err error) {
	err = repo.db.Delete(&m.Major{}, id).Error
	return
}

func (repo majorRepository) ReadMajor() (major *[]m.Major, err error) {
	err = repo.db.Find(&major).Error
	return
}
