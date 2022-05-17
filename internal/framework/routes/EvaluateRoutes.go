package routes

import (
	"go-question-board/internal/framework/transport/controller"
	 mw "go-question-board/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewEvaluateRoutes(e *echo.Group, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	newTeacherEvaluateRoutes(e.Group("/teacher"), ucon)
	newStudentEvaluateRoutes(e.Group("/student"), ucon)
}

func newTeacherEvaluateRoutes(e *echo.Group, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.TeacherPermission)
	group := e.Group("/evaluate", middleware...)
	group.GET("", ucon.ViewEvaluateResponse)
}

func newStudentEvaluateRoutes(e *echo.Group, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.StudentPermission)
	group := e.Group("/evaluate", middleware...)
	group.GET("", ucon.GetQuest)
	group.POST("", ucon.QuestAnswer)
}
