package repository

import (
	m "go-question-board/internal/core/models"
)

type UserRepository interface {
	CreateUser(m.UserModel) error
	UpdateUser(int, m.UserModel) error
	DeleteUser(int) error
	ReadUser() (*[]m.UserModel, error)
	ReadUserByID(int) (*m.UserModel, error)
}
