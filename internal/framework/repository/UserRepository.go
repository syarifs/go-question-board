package repository

import (
	m "go-question-board/internal/core/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo userRepository) CreateUser(user m.UserModel) (err error) {
	err = repo.db.Omit("Major", "Level").Create(&user).Error
	return
}

func (repo userRepository) UpdateUser(id int, user m.UserModel) (err error) {
	err = repo.db.Where("id = ?", id).Updates(&user).Error
	return
}

func (repo userRepository) DeleteUser(id int) (err error) {
	err = repo.db.Delete(&m.UserModel{}, id).Error
	return
}

func (repo userRepository) ReadUser() (user *[]m.UserModel, err error) {
	err = repo.db.Preload(clause.Associations).Find(&user).Error
	return
}

func (repo userRepository) ReadUserByID(id int) (user *m.UserModel, err error) {
	err = repo.db.Preload(clause.Associations).Find(&user, id).Error
	return
}
