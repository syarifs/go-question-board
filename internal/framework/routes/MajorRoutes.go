package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewMajorRoutes(e *echo.Echo, ucon *controller.MajorController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/major", middleware...)
	group.GET("", ucon.ReadMajor)
	group.POST("", ucon.CreateMajor)
	group.PUT("/:id/update", ucon.UpdateMajor)
	group.DELETE("/:id/delete", ucon.DeleteMajor)
}
