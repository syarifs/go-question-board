package repository

import (
	"go-question-board/internal/core/models"
	req "go-question-board/internal/core/models/request"
)

type AuthRepository interface {
	Login(req.LoginRequest) (models.User, error)
	RefreshToken(models.Token) (models.Token, error)
}
