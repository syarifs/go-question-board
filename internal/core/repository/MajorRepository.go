package repository

import (
	m "go-question-board/internal/core/models"
)

type MajorRepository interface {
	CreateMajor(m.Major) error
	UpdateMajor(m.Major) error
	DeleteMajor(int) error
	ReadMajor() (*[]m.Major, error)
}
