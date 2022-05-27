package controller

import (
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/service"
	"go-question-board/internal/utils/errors"
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

// CreateResource godoc
// @Summary Get Evaluate Quest
// @Description Route Path for Get List of Evaluation Quest with Subject ID and Class.
// @Tags Evaluate
// @Security ApiKey
// @Param body body request.User{} true "send user data"
// @Param subject_id query int true "subject id"
// @Success 200 {object} response.MessageData{data=response.Quest{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /student/evaluate [get]
func (ucon EvaluateController) GetQuest(c echo.Context) error {
	var user request.User
	c.Bind(&user)

	subject_id, _ := strconv.Atoi(c.QueryParam("subject_id"))

	res, err := ucon.srv.GetQuest(subject_id, user)

	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Quest Fetched",
			Data: res,
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Fetch Quest",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Answer Evaluation Quest
// @Description Route Path for Answer Evaluation Quest.
// @Tags Evaluate
// @Security ApiKey
// @Accept json
// @Produce json
// @Param class query string false "class"
// @Param teacher_id query int false "teacher id"
// @Param subject_id query int false "subject id"
// @Param body body request.Answer{} true "send quest data"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /student/evaluate [post]
func (ucon EvaluateController) QuestAnswer(c echo.Context) error {
	questAnswer := request.Answer{}
	class := c.QueryParam("class")
	teacher_id, _ := strconv.Atoi(c.QueryParam("teacher_id"))
	subject_id, _ := strconv.Atoi(c.QueryParam("subject_id"))

	c.Bind(&questAnswer)

	err := ucon.srv.Evaluate(questAnswer, teacher_id, subject_id, class)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Quest Answered",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Answer Quest",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Quest Response By Quest ID
// @Description Route Path for Get Quest Response By Quest ID.
// @Tags Evaluate
// @Security ApiKey
// @Param subject_id query int true "subject id"
// @Param class query int true "class"
// @Success 200 {object} response.MessageData{data=response.QuestResponses} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /teacher/evaluate [get]
func (ucon EvaluateController) ViewEvaluateResponse(c echo.Context) error {
	var user response.User
	c.Bind(&user)
	class := c.QueryParam("class")
	subject_id, _ := strconv.Atoi(c.QueryParam("subject_id"))
	res, err := ucon.srv.ViewEvaluateResponse(int(user.ID), subject_id, class)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Quest Response Fetched",
			Data: res,
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Fetch Quest Response",
			Error: err.Error(),
		})
	}
}

