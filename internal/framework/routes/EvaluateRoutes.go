package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewEvaluateRoutes(e *echo.Group, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	newTeacherEvaluateRoutes(e.Group("/teacher"), ucon)
	newStudentEvaluateRoutes(e.Group("/student"), ucon)
}

func newTeacherEvaluateRoutes(e *echo.Group, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/evaluate", middleware...)
	group.GET("", ucon.ViewEvaluateResponse)
}

func newStudentEvaluateRoutes(e *echo.Group, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/evaluate", middleware...)
	group.GET("", ucon.GetQuest)
	group.POST("", ucon.QuestAnswer)
}