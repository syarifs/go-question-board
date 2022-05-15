package repository

import (
	"go-question-board/internal/core/models"
)

type AuthRepository interface {
	Login(string) (models.User, error)
	RefreshToken(models.Token) (models.Token, error)
}
