package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewRoutes (e *echo.Group, ctrl *controller.Controller) {
	NewAuthRoutes(e, ctrl.Auth)
	NewUserRoutes(e, ctrl.User)
	NewMajorRoutes(e, ctrl.Major)
	NewTagRoutes(e, ctrl.Tag)
	NewSubjectRoutes(e, ctrl.Subject)
	NewQuestionnaireRoutes(e, ctrl.Questionnare)
	NewEvaluateRoutes(e, ctrl.Evaluate)
}
