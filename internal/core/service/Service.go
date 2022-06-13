package service

import "go-question-board/internal/core/repository"


type Service struct {
	Auth *AuthService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(r.Auth),
	}
}
