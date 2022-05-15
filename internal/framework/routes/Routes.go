package routes

import (
	"go-question-board/internal/framework/transport/controller"
	// "go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewRoutes (e *echo.Echo, ctrl *controller.Controller) {
	NewAuthRoutes(e, ctrl.Auth)
	NewUserRoutes(e, ctrl.User)
	NewMajorRoutes(e, ctrl.Major)
	NewTagRoutes(e, ctrl.Tag)
	NewSubjectRoutes(e, ctrl.Subject)
	NewQuestionnaireRoutes(e, ctrl.Questionnare)
	NewEvaluateRoutes(e, ctrl.Evaluate)
}
