package controller

import (
	"fmt"
	"go-question-board/internal/core/models"
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
// @Summary Create New Questionnaire
// @Description Route Path for Insert New Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  models.Questionnaire{}  true "send request questionnaire code and questionnaire name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /questionnaire [post]
func (ucon QuestionnaireController) CreateQuestionnaire(c echo.Context) error {
	questionnaire := models.Questionnaire{}
	c.Bind(&questionnaire)
	err := ucon.srv.CreateQuest(questionnaire)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Questionnaire Created",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Create Questionnaire",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Questionnaire
// @Description Route Path for Get List of Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /questionnaire/{id} [get]
func (ucon QuestionnaireController) ViewQuestByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ucon.srv.ViewQuestByID(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Questionnaire Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Questionnaire",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Questionnaire
// @Description Route Path for Get List of Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /questionnaire/{id}/response [get]
func (ucon QuestionnaireController) ViewQuestResponse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ucon.srv.ViewQuestResponse(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Questionnaire Response Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Questionnaire Response",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Questionnaire
// @Description Route Path for Get List of Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageData{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /questionnaire [get]
func (ucon QuestionnaireController) ListMyQuestionnaire(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	fmt.Println(user.ID)
	res, err := ucon.srv.MyQuest(int(user.ID))
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Questionnaire Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Questionnaire",
			Error: err,
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
// @Param id path int true "questionnaire id"
// @Param body  body  models.Questionnaire{}  true "send request questionnaire code and questionnaire name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /questionnaire/{id}/update [PUT]
func (ucon QuestionnaireController) UpdateQuestionnaire(c echo.Context) error {
	questionnaire := models.Questionnaire{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&questionnaire)
	err := ucon.srv.UpdateQuest(id, questionnaire)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Questionnaire Updated",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Update Questionnaire",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Delete Questionnaire
// @Description Route Path for Delete Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "questionnaire id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /questionnaire/{id}/delete [DELETE]
func (ucon QuestionnaireController) DeleteQuestionnaire(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteQuest(id)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Questionnaire Deleted",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Delete Questionnaire",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Questionnaire
// @Description Route Path for Get List of Questionnaire.
// @Tags Questionnaire
// @Security ApiKey
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Router /dashboard [get]
func (ucon QuestionnaireController) AvailableQuest(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	res, err := ucon.srv.QuestForMe(user.Tags)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Questionnaire Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Questionnaire",
			Error: err,
		})
	}
}

