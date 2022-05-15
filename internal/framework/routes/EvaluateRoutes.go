package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewEvaluateRoutes(e *echo.Echo, ucon *controller.EvaluateController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/evaluate", middleware...)
	group.GET("", ucon.GetQuest)
	group.POST("", ucon.QuestAnswer)
}
