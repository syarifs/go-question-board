package controller

import "go-question-board/internal/core/service"

type Controller struct {
	Auth *AuthController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
	}
}
