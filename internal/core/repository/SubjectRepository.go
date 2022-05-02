package repository

import (
	m "go-question-board/internal/core/models"
)

type SubjectRepository interface {
	CreateSubject(m.Subject) error
	UpdateSubject(int, m.Subject) error
	DeleteSubject(int) error
	ReadSubject() (*[]m.Subject, error)
}
