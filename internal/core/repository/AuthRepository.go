package repository

import "go-question-board/internal/core/entity/models"

type AuthRepository interface {
	Login(string) (models.User, error)
	RefreshToken(models.Token) (models.Token, error)
}
