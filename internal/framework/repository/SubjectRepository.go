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

func (repo subjectRepository) CreateSubject(subject m.SubjectModel) (err error) {
	err = repo.db.Save(&subject).Error
	return
}

func (repo subjectRepository) UpdateSubject(id int, subject m.SubjectModel) (err error) {
	err = repo.db.Where("id = ?", id).Updates(&subject).Error
	return
}

func (repo subjectRepository) DeleteSubject(id int) (err error) {
	err = repo.db.Delete(&m.SubjectModel{}, id).Error
	return
}

func (repo subjectRepository) ReadSubject() (subject *[]m.SubjectModel, err error) {
	err = repo.db.Find(&subject).Error
	return
}
