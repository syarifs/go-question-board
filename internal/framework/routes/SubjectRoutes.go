package routes

import (
	"go-question-board/internal/framework/transport/controller"
	// mw "go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewSubjectRoutes(e *echo.Group, ucon *controller.SubjectController) {
	newAdminSubjectRoutes(e.Group("/admin"), ucon)
	newTeacherSubjectRoutes(e.Group("/teacher"), ucon)
	newStudentSubjectRoutes(e.Group("/student"), ucon)
}

func newAdminSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadSubject)
	group.GET("/:id", ucon.ReadSubjectByID)
	group.POST("", ucon.CreateSubject)
	group.PUT("/:id/update", ucon.UpdateSubject)
	group.DELETE("/:id/delete", ucon.DeleteSubject)
}

func newTeacherSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadTeacherSubject)
	group.POST("", ucon.CreateSubject)
}

func newStudentSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadStudentSubject)
	group.POST("", ucon.CreateSubject)
}
