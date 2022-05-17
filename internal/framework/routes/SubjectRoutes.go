package routes

import (
	"go-question-board/internal/framework/transport/controller"
	mw "go-question-board/internal/framework/transport/middleware"
	// mw "go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	newAdminSubjectRoutes(e.Group("/admin"), ucon, middleware...)
	newTeacherSubjectRoutes(e.Group("/teacher"), ucon, middleware...)
	newStudentSubjectRoutes(e.Group("/student"), ucon, middleware...)
}

func newAdminSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.AdminPermission)
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadSubject)
	group.GET("/:id", ucon.ReadSubjectByID)
	group.POST("", ucon.CreateSubject)
	group.PUT("/:id/update", ucon.UpdateSubject)
	group.DELETE("/:id/delete", ucon.DeleteSubject)
}

func newTeacherSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.TeacherPermission)
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadTeacherSubject)
	group.POST("", ucon.CreateSubject)
}

func newStudentSubjectRoutes(e *echo.Group, ucon *controller.SubjectController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.StudentPermission)
	group := e.Group("/subject", middleware...)
	group.GET("", ucon.ReadStudentSubject)
	group.POST("", ucon.CreateSubject)
}
