package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, ucon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	e.GET("/user", ucon.ReadUser, middleware...)
	e.POST("/user", ucon.CreateUser, middleware...)
	e.PUT("/user/:id/update", ucon.UpdateUser, middleware...)
	e.DELETE("/user/:id/delete", ucon.DeleteUser, middleware...)
}
