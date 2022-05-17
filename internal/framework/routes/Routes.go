package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewRoutes (e *echo.Group, ctrl *controller.Controller, middleware ...echo.MiddlewareFunc) {
	NewAuthRoutes(e, ctrl.Auth)
	NewUserRoutes(e, ctrl.User, middleware...)
	NewMajorRoutes(e, ctrl.Major, middleware...)
	NewTagRoutes(e, ctrl.Tag, middleware...)
	NewSubjectRoutes(e, ctrl.Subject, middleware...)
	NewQuestionnaireRoutes(e, ctrl.Questionnare, middleware...)
	NewEvaluateRoutes(e, ctrl.Evaluate, middleware...)
}
