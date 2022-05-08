package repository

import (
	m "go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/framework/transport/middleware"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db: db}
}

func (repo authRepository) Login(user request.LoginRequest) (users m.User, err error) {
	err = repo.db.Preload(clause.Associations).Where("email = ?", user.Email).First(&users).Error
	return
}

func (repo authRepository) RefreshToken(str m.Token) (token m.Token, err error) {
	token, err = middleware.RefreshToken(str)
	return
}
