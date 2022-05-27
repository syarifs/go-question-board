package repository

import (
	m "go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *subjectRepository {
	return &subjectRepository{db: db}
}

func (repo subjectRepository) CreateSubject(subject m.Subject) (err error) {
	err = repo.db.Create(&subject).Error
	return
}

func (repo subjectRepository) UpdateSubject(subject m.Subject) (err error) {
	err = repo.db.Updates(&subject).Error
	return
}

func (repo subjectRepository) DeleteSubject(id int) (err error) {
	del := repo.db.Delete(&m.Subject{}, id)
	if del.RowsAffected < 1 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (repo subjectRepository) ReadSubject() (subject *[]m.Subject, err error) {
	err = repo.db.Preload("Major").Find(&subject).Error
	return
}

func (repo subjectRepository)	ReadTeacherSubject(id int) (sub *[]m.Subject, err error) {
	var subject_id []int

	repo.db.Table("teacher_subjects").Select("subject_id").
			Where("user_id = ?", id).Scan(&subject_id)

	err = repo.db.
		Preload("Major").
		Preload("Student").
		Preload("Student.Tags").
		Preload("Student.Major").
		Where("id IN (?)", subject_id).
		Find(&sub).Error
	return
}

func (repo subjectRepository) ReadStudentSubject(id int) (sub *[]m.Subject, err error) {
	var class string

	repo.db.Model(m.Tag{}).Select("value").
		Where("id IN (?) AND name = 'Class'", repo.db.Table("user_tags").
			Where("user_id = ?", id).Select("tag_id")).Scan(&class)

	err = repo.db.
		Preload("Major").
		Preload("Teacher", "class = ?", class).
		Preload("Teacher.User").
		Find(&sub, "id IN (?)", repo.db.
			Table("student_subject").
			Select("subject_id").Where("user_id = ?", id)).Error
	return
}

func (repo subjectRepository) ReadSubjectByID(id int) (subject *m.Subject, err error) {
	err = repo.db.Debug().
		Preload("Major").
		Preload("Student.Major").
		Preload("Student.Tags").
		Preload("Teacher").
		Preload("Teacher.User").
		First(&subject, id).Error
	return
}
