package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewQuestionnaireRoutes(e *echo.Group, ucon *controller.QuestionnaireController, middleware ...echo.MiddlewareFunc) {
	group := e.Group("/quest", middleware...)
	group.GET("", ucon.MyQuest)
	group.GET("/available", ucon.QuestForMe)
	group.GET("/:id", ucon.ViewQuestByID)
	group.GET("/:id/response", ucon.ViewQuestResponse)
	group.POST("", ucon.CreateQuest)
	group.POST("/answer", ucon.QuestAnswer)
	group.PUT("/:id/update", ucon.UpdateQuest)
	group.DELETE("/:id/delete", ucon.DeleteQuest)
}
