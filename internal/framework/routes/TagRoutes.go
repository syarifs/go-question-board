package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewTagRoutes(e *echo.Echo, ucon *controller.TagController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/tag", middleware...)
	group.GET("", ucon.ReadTag)
	group.POST("", ucon.CreateTag)
	group.PUT("/:id/update", ucon.UpdateTag)
	group.DELETE("/:id/delete", ucon.DeleteTag)
}
