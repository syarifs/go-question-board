package repository

import (
	m "go-question-board/internal/core/entity/models"
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
	err = repo.db.Omit("Tags.*", "Subject.*").Create(&user).Error
	return
}

func (repo userRepository) UpdateUser(user m.User) (err error) {
	if ! utils.IsEmpty(user.TeacherSubject) {
		repo.db.Model(m.TeacherSubject{}).Delete("user_id = ?", user.ID)
		repo.db.Model(m.TeacherSubject{}).Create(&user.TeacherSubject)
	}

	err = repo.db.Omit(clause.Associations).Updates(&user).Error

	if err == nil {
		err = repo.db.Model(&user).Association("Tags").Replace(&user.Tags)
		err = repo.db.Model(&user).Association("Subject").Replace(&user.Subject)
	}

	return
}

func (repo userRepository) DeleteUser(id int) (err error) {
	del := repo.db.Delete(&m.User{}, id)
	if del.RowsAffected < 1 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (repo userRepository) ReadUser() (user *[]m.User, err error) {
	err = repo.db.Preload(clause.Associations).
		Find(&user).Error
	return
}

func (repo userRepository) ReadUserByID(id int) (user *m.User, err error) {
	var class string

	repo.db.Model(m.Tag{}).Select("value").
		Where("id IN (?) AND name = 'Class'", repo.db.Table("user_tags").
			Where("user_id = ?", id).Select("tag_id")).Scan(&class)

	err = repo.db.Preload(clause.Associations).
		Preload("TeacherSubject.Subject").
		Preload("TeacherSubject.Subject.Major").
		Preload("TeacherSubject.User").
		Preload("Subject").
		Preload("Subject.Major").
		Preload("Subject.Teacher.User").
		Preload("Subject.Teacher", "class = ?", class).
		First(&user, id).Error
	return
}
