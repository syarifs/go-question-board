package repository

import (
	m "go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (repo authRepository) Login(email string) (users m.User, err error) {
	err = repo.db.
		Preload(clause.Associations).
		Where("email = ?", email).First(&users).Error
	return
}

func (repo authRepository) SaveToken(token m.Token) (err error) {
	return
}

func (repo authRepository) RefreshToken(str m.Token) (token m.Token, err error) {
	return
}
