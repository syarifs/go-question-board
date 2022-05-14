package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/service"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/labstack/echo/v4"
)

type QuestionnaireController struct {
	srv *service.QuestionnaireService
}

func NewQuestionnaireController(srv *service.QuestionnaireService) *QuestionnaireController {
	return &QuestionnaireController{srv}
}

// CreateResource godoc
// @Summary Create New Quest
// @Description Route Path for Insert New Quest.
// @Tags Questionnaire
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  models.Questionnaire{}  true "send quest data"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest [post]
func (ucon QuestionnaireController) CreateQuest(c echo.Context) error {
	questionnaire := models.Questionnaire{}
	c.Bind(&questionnaire)
	err := ucon.srv.CreateQuest(questionnaire)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Quest Created",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Create Quest",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Quest
// @Description Route Path for Get List of Quest.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest/{id} [get]
func (ucon QuestionnaireController) ViewQuestByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ucon.srv.ViewQuestByID(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Quest Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Quest",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Quest
// @Description Route Path for Get List of Quest.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageData{data=response.QuestResponses} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest/{id}/response [get]
func (ucon QuestionnaireController) ViewQuestResponse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ucon.srv.ViewQuestResponse(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Quest Response Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Quest Response",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Quest
// @Description Route Path for Get List of Quest By User ID.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageData{data=[]response.QuestList{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest [get]
func (ucon QuestionnaireController) MyQuest(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	res, err := ucon.srv.MyQuest(int(user.ID))
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

// CreateResource godoc
// @Summary Update Questionnaire
// @Description Route Path for Update Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "quest id"
// @Param body  body  models.Questionnaire{}  true "send quest data"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest/{id}/update [PUT]
func (ucon QuestionnaireController) UpdateQuest(c echo.Context) error {
	questionnaire := models.Questionnaire{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&questionnaire)
	err := ucon.srv.UpdateQuest(id, questionnaire)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Quest Updated",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Update Quest",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Delete Quest
// @Description Route Path for Delete Quest.
// @Tags Questionnaire
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "quest id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest/{id}/delete [DELETE]
func (ucon QuestionnaireController) DeleteQuest(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteQuest(id)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Quest Deleted",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Delete Quest",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Get All Quest with Tag Filter
// @Description Route Path for Get List of Quest with User Tag Filter.
// @Tags Questionnaire
// @Security ApiKey
// @Param body body models.User{} true "send logged in user data"
// @Success 200 {object} response.MessageData{data=[]response.AvailableQuestList} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Router /quest/available [get]
func (ucon QuestionnaireController) QuestForMe(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	res, err := ucon.srv.QuestForMe(int(user.ID), user.Tags)
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

// CreateResource godoc
// @Summary Answer Quest
// @Description Route Path for Answer Quest.
// @Tags Questionnaire
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  request.Answer{}  true "send answer, quest, and user data"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /quest/answer [post]
func (ucon QuestionnaireController) QuestAnswer(c echo.Context) error {
	questAnswer := request.Answer{}
	c.Bind(&questAnswer)
	err := ucon.srv.AnswerQuest(questAnswer)
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
