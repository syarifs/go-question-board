package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewMajorRoutes(e *echo.Echo, ucon *controller.MajorController, middleware ...echo.MiddlewareFunc) {
	e.GET("/major", ucon.ReadMajor, middleware...)
	e.POST("/major", ucon.CreateMajor, middleware...)
	e.PUT("/major/:id/update", ucon.UpdateMajor, middleware...)
	e.DELETE("/major/:id/delete", ucon.DeleteMajor, middleware...)
}
