package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewSubjectRoutes(e *echo.Echo, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	e.GET("/subject", ucon.ReadSubject, middleware...)
	e.POST("/subject", ucon.CreateSubject, middleware...)
	e.PUT("/subject/:id/update", ucon.UpdateSubject, middleware...)
	e.DELETE("/subject/:id/delete", ucon.DeleteSubject, middleware...)
}
