package repository

import m "go-question-board/internal/core/entity/models"

type SubjectRepository interface {
	CreateSubject(m.Subject) error
	UpdateSubject(m.Subject) error
	DeleteSubject(int) error
	ReadSubject() (*[]m.Subject, error)
	ReadSubjectByID(int) (*m.Subject, error)
	ReadTeacherSubject(int) (*[]m.Subject, error)
	ReadStudentSubject(int) (*[]m.Subject, error)
}
