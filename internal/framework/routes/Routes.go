package routes

import (
	"go-question-board/internal/framework/transport/controller"
	"go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewRoutes (e *echo.Echo, ctrl *controller.Controller) {
	NewAuthRoutes(e, ctrl.Auth)
	NewUserRoutes(e, ctrl.User, middleware.JWT(), middleware.AdminPermission)
	NewMajorRoutes(e, ctrl.Major, middleware.JWT(), middleware.AdminPermission)
	NewTagRoutes(e, ctrl.Tag, middleware.JWT(), middleware.AdminPermission)
	NewSubjectRoutes(e, ctrl.Subject, middleware.JWT(), middleware.AdminPermission)
}
