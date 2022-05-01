package repository

import (
	m "go-question-board/internal/core/models"
)

type SubjectRepository interface {
	CreateSubject(m.SubjectModel) error
	UpdateSubject(int, m.SubjectModel) error
	DeleteSubject(int) error
	ReadSubject() (*[]m.SubjectModel, error)
}
