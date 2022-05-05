package repository

import (
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
)

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *subjectRepository {
	return &subjectRepository{db: db}
}

func (repo subjectRepository) CreateSubject(subject m.Subject) (err error) {
	err = repo.db.Save(&subject).Error
	return
}

func (repo subjectRepository) UpdateSubject(subject m.Subject) (err error) {
	err = repo.db.Updates(&subject).Error
	return
}


func (repo subjectRepository) DeleteSubject(id int) (err error) {
	err = repo.db.Delete(&m.Subject{}, id).Error
	return
}

func (repo subjectRepository) ReadSubject() (subject *[]m.Subject, err error) {
	err = repo.db.Find(&subject).Error
	return
}
