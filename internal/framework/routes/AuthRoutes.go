package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewAuthRoutes(e *echo.Echo, acon *controller.AuthController, middleware ...echo.MiddlewareFunc) {
	e.POST("/refresh_token", acon.RefreshToken, middleware...)
	e.POST("/login", acon.Login, middleware...)
}
