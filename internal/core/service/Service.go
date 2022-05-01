package service

import "go-question-board/internal/core/repository"

type Service struct {
	Auth *AuthService
	User *UserService
	Tag *TagService
	Major *MajorService
	Subject *SubjectService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(r.Auth),
		User: NewUserService(r.User),
		Tag: NewTagService(r.Tag),
		Major: NewMajorService(r.Major),
		Subject: NewSubjectService(r.Subject),
	}
}
