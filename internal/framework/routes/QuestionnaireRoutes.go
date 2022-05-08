package routes

import (
	"go-question-board/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewQuestionnaireRoutes(e *echo.Echo, ucon *controller.QuestionnaireController, middleware ...echo.MiddlewareFunc) {
	e.GET("/dashboard", ucon.AvailableQuest, middleware...)
	group := e.Group("/questionnaire", middleware...)
	group.GET("", ucon.ListMyQuestionnaire)
	group.GET("/:id", ucon.ViewQuestByID)
	group.POST("", ucon.CreateQuestionnaire)
	group.PUT("/:id/update", ucon.UpdateQuestionnaire)
	group.DELETE("/:id/delete", ucon.DeleteQuestionnaire)
}
