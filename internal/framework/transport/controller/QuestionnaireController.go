package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/service"
	"net/http"

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
// @Description Route Path for Insert New Questionnaire, for Administrator only.
// @Tags Questionnaire
// @Accept json
// @Produce json
// @Param data  body  models.Questionnaire{}  true "send request questionnaire code and questionnaire name"
// @Success 200 {object} response.SuccessResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Router /questionnaire [post]
func (ucon QuestionnaireController) CreateQuestionnaire(c echo.Context) error {
	questionnaire := models.Questionnaire{}
	c.Bind(&questionnaire)
	res, err := ucon.srv.CreateQuestionnaire(questionnaire)
	if err == nil {
		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Message: "Questionnaire Created",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
			Message: "Failed to Create Questionnaire",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Questionnaire
// @Description Route Path for Get List of Questionnaire, for Administrator only.
// @Tags Questionnaire
// @Success 200 {object} response.SuccessResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Failure 400 {object} string error
// @Router /questionnaire [get]
func (ucon QuestionnaireController) ListMyQuestionnaire(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	res, err := ucon.srv.MyQuestionnaire(int(user.ID))
	if err == nil {
		return c.JSON(http.StatusOK, response.SuccessResponse{
			Message: "Questionnaire Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
			Message: "Failed to Fetch Questionnaire",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Update Questionnaire
// @Description Route Path for Update Questionnaire, for Administrator only.
// @Tags Questionnaire
// @Accept json
// @Produce json
// @Param data  body  models.Questionnaire{}  true "send request questionnaire code and questionnaire name"
// @Param id path int true "questionnaire id"
// @Success 200 {object} response.SuccessResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Router /questionnaire/{id}/update [PUT]
// func (ucon QuestionnaireController) UpdateQuestionnaire(c echo.Context) error {
// 	questionnaire := models.Questionnaire{}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	c.Bind(&questionnaire)
// 	res, err := ucon.srv.UpdateQuestionnaire(id, questionnaire)
// 	if err == nil {
// 		return c.JSON(http.StatusCreated, response.SuccessResponse{
// 			Message: "Questionnaire Updated",
// 			Data: res,
// 		})
// 	} else {
// 		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
// 			Message: "Failed to Update Questionnaire",
// 			Error: err,
// 		})
// 	}
// }

// CreateResource godoc
// @Summary Delete Questionnaire
// @Description Route Path for Delete Questionnaire, for Administrator only.
// @Tags Questionnaire
// @Accept json
// @Produce json
// @Param id path int true "questionnaire id"
// @Success 200 {object} response.MessageOnlyResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Router /questionnaire/{id}/delete [DELETE]
// func (ucon QuestionnaireController) DeleteQuestionnaire(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	err := ucon.srv.DeleteQuestionnaire(id)
// 	if err == nil {
// 		return c.JSON(http.StatusCreated, response.MessageOnlyResponse{
// 			Message: "Questionnaire Updated",
// 		})
// 	} else {
// 		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
// 			Message: "Failed to Update Questionnaire",
// 			Error: err,
// 		})
// 	}
// }
