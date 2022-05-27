package repository

import "go-question-board/internal/core/entity/models"

type MajorRepository interface {
	CreateMajor(models.Major) error
	UpdateMajor(models.Major) error
	DeleteMajor(int) error
	ReadMajor() (*[]models.Major, error)
}
