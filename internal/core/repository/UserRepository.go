package repository

import (
	m "go-question-board/internal/core/models"
)

type UserRepository interface {
	CreateUser(m.User) error
	UpdateUser(int, m.User) error
	DeleteUser(int) error
	ReadUser() (*[]m.User, error)
	ReadUserByID(int) (*m.User, error)
}
