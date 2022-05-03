package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewSubjectRoutes(e *echo.Echo, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadSubject)
	group.POST("", ucon.CreateSubject)
	group.PUT("/:id/update", ucon.UpdateSubject)
	group.DELETE("/:id/delete", ucon.DeleteSubject)
}
