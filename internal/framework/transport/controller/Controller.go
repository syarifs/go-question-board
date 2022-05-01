package controller

import "go-question-board/internal/core/service"

type Controller struct {
	Auth *AuthController
	User *UserController
	Major *MajorController
	Tag *TagController
	Subject *SubjectController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
		User: NewUserController(srv.User),
		Major: NewMajorController(srv.Major),
		Subject: NewSubjectController(srv.Subject),
		Tag: NewTagController(srv.Tag),
	}
}
