package repository

import (
	m "go-question-board/internal/core/models"
	"go-question-board/internal/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo userRepository) CreateUser(user m.User) (err error) {
	user.Password = utils.HashPassword(user.Password)
	err = repo.db.Omit("Tags.*", "Subject.*").Create(&user).Error
	return
}

func (repo userRepository) UpdateUser(user m.User) (err error) {
	err = repo.db.Updates(&user).Error
	if err == nil {
		err = repo.db.Model(&user).Association("Tags").Replace(&user.Tags)
		err = repo.db.Model(&user).Association("Subject").Replace(&user.Subject)
	}
	return
}

func (repo userRepository) DeleteUser(id int) (err error) {
	err = repo.db.Delete(&m.User{}, id).Error
	return
}

func (repo userRepository) ReadUser() (user *[]m.User, err error) {
	err = repo.db.Preload(clause.Associations).
		Find(&user).Error
	return
}

func (repo userRepository) ReadUserByID(id int) (user *m.User, err error) {
	var class string

	repo.db.Find(m.Tag{}, "id = ? AND name = 'class'", repo.db.Table("user_tags").
		Select("subject_id").
		Where("user_id = ?", id)).Scan(&class)

	err = repo.db.Debug().Preload(clause.Associations).
		Preload("TeacherSubject.Subject").
		Preload("TeacherSubject.Subject.Major").
		Preload("TeacherSubject.User").
		Preload("Subject").
		Preload("Subject.Major").
		Preload("Subject.Teacher.User").
		Preload("Subject.Teacher", "class = ?", class).
		Find(&user, id).Error
	return
}
