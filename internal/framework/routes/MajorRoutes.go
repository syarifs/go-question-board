package routes

import (
	"go-question-board/internal/framework/transport/controller"
	mw "go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewMajorRoutes(e *echo.Group, ucon *controller.MajorController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.AdminPermission)
	group := e.Group("/admin/major", middleware...)
	group.GET("", ucon.ReadMajor)
	group.POST("", ucon.CreateMajor)
	group.PUT("/:id/update", ucon.UpdateMajor)
	group.DELETE("/:id/delete", ucon.DeleteMajor)
}
