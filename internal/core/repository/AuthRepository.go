package repository

import (
	"go-question-board/internal/core/models"
	req "go-question-board/internal/core/models/request"
)

type AuthRepository interface {
	Login(req.LoginRequest) (models.UserModel, error)
	RefreshToken(models.TokenModel) (models.TokenModel, error)
}
