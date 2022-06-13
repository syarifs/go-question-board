package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewRoutes (e *echo.Group, ctrl *controller.Controller, middleware ...echo.MiddlewareFunc) {
	NewAuthRoutes(e, ctrl.Auth)
}
