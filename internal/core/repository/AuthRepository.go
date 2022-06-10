package repository

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/response"
)

type AuthRepository interface {
	Login(string) (response.UserDetails, error)
	SaveToken(models.Token) (error)
	UpdateToken(models.Token, models.Token) (error)
	RevokeToken(string) (error)
}
