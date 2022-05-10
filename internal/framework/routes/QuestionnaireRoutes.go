package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewQuestionnaireRoutes(e *echo.Echo, ucon *controller.QuestionnaireController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/quest", middleware...)
	group.GET("", ucon.ListMyQuestionnaire)
	group.GET("/available", ucon.AvailableQuest)
	group.GET("/:id", ucon.ViewQuestByID)
	group.GET("/:id/response", ucon.ViewQuestResponse)
	group.POST("", ucon.CreateQuestionnaire)
	group.PUT("/:id/update", ucon.UpdateQuestionnaire)
	group.DELETE("/:id/delete", ucon.DeleteQuestionnaire)
}
