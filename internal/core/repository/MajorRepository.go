package repository

import (
	m "go-question-board/internal/core/models"
)

type MajorRepository interface {
	CreateMajor(m.MajorModel) error
	UpdateMajor(int, m.MajorModel) error
	DeleteMajor(int) error
	ReadMajor() (*[]m.MajorModel, error)
}
