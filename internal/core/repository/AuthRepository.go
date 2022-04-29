package repository

import (
	req "go-question-board/internal/core/models/request"
	res "go-question-board/internal/core/models/response"
)

type AuthRepository interface {
	CreateLevel(req.LevelRequest) error
	Login(req.LoginRequest) (res.UserResponse, error)
}
