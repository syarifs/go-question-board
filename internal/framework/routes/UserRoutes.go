package routes

import (
	"go-question-board/internal/framework/transport/controller"
	mw "go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Group, ucon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	newAdminUserRoutes(e.Group("/admin"), ucon, middleware...)
}

func newAdminUserRoutes(e *echo.Group, ucon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.AdminPermission)
	group := e.Group("/user", middleware...)
	group.GET("", ucon.ReadUser)
	group.POST("", ucon.CreateUser)
	group.GET("/:id", ucon.ReadUserByID)
	group.PUT("/:id/update", ucon.UpdateUser)
	group.DELETE("/:id/delete", ucon.DeleteUser)
}

func NewProfileRoutes(e *echo.Group, ucon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/profile", middleware...)
	group.GET("", ucon.ReadUser)
	group.PUT("/update", ucon.UpdateUser)
}
