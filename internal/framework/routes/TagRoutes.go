package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewTagRoutes(e *echo.Echo, ucon *controller.TagController, middleware ...echo.MiddlewareFunc) {
	e.GET("/tag", ucon.ReadTag, middleware...)
	e.POST("/tag", ucon.CreateTag, middleware...)
	e.PUT("/tag/:id/update", ucon.UpdateTag, middleware...)
	e.DELETE("/tag/:id/delete", ucon.DeleteTag, middleware...)
}
