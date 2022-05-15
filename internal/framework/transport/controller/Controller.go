package controller

import "go-question-board/internal/core/service"

type Controller struct {
	Auth *AuthController
	User *UserController
	Major *MajorController
	Tag *TagController
	Subject *SubjectController
	Questionnare *QuestionnaireController
	Evaluate *EvaluateController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
		User: NewUserController(srv.User),
		Major: NewMajorController(srv.Major),
		Subject: NewSubjectController(srv.Subject),
		Tag: NewTagController(srv.Tag),
		Questionnare: NewQuestionnaireController(srv.Question),
		Evaluate: NewEvaluateController(srv.Evaluate),
	}
}
