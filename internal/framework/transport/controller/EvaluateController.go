package controller

import (
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EvaluateController struct {
	srv *service.EvaluateService
}

func NewEvaluateController(srv *service.EvaluateService) *EvaluateController {
	return &EvaluateController{srv}
}

func (ucon EvaluateController) GetQuest(c echo.Context) error {
	res, err := ucon.srv.GetQuest()
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Quest Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Quest",
			Error: err.Error(),
		})
	}
}

func (ucon EvaluateController) QuestAnswer(c echo.Context) error {
	questAnswer := request.Answer{}
	teacher_id, _ := strconv.Atoi(c.QueryParam("teacher_id"))
	subject_id, _ := strconv.Atoi(c.QueryParam("subject_id"))

	c.Bind(&questAnswer)

	err := ucon.srv.Evaluate(questAnswer, teacher_id, subject_id)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Quest Answered",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Answer Quest",
			Error: err.Error(),
		})
	}
}
