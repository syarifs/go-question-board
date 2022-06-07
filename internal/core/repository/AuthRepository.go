package repository

import "go-question-board/internal/core/entity/models"

type AuthRepository interface {
	Login(string) (models.User, error)
	SaveToken(models.Token) (error)
	UpdateToken(models.Token, models.Token) (error)
	RevokeToken(models.Token) (error)
}
