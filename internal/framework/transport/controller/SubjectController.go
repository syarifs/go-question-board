package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/service"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type SubjectController struct {
	srv *service.SubjectService
}

func NewSubjectController(srv *service.SubjectService) *SubjectController {
	return &SubjectController{srv}
}

// CreateResource godoc
// @Summary Create New Subject
// @Description Route Path for Insert New Subject, for Administrator only.
// @Tags Subject
// @Accept json
// @Produce json
// @Param data  body  models.Subject{}  true "send request subject code and subject name"
// @Success 200 {object} response.SuccessResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Router /subject [post]
func (ucon SubjectController) CreateSubject(c echo.Context) error {
	subject := models.Subject{}
	c.Bind(&subject)
	res, err := ucon.srv.CreateSubject(subject)
	if err == nil {
		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Message: "Subject Created",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
			Message: "Failed to Create Subject",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All Subject
// @Description Route Path for Get List of Subject, for Administrator only.
// @Tags Subject
// @Success 200 {object} response.SuccessResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Failure 400 {object} string error
// @Router /subject [get]
func (ucon SubjectController) ReadSubject(c echo.Context) error {
	res, err := ucon.srv.ReadSubject()
	if err == nil {
		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Message: "Subject Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
			Message: "Failed to Fetch Subject",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Update Subject
// @Description Route Path for Update Subject, for Administrator only.
// @Tags Subject
// @Accept json
// @Produce json
// @Param data  body  models.Subject{}  true "send request subject code and subject name"
// @Param id path int true "subject id"
// @Success 200 {object} response.SuccessResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Router /subject/{id}/update [PUT]
func (ucon SubjectController) UpdateSubject(c echo.Context) error {
	subject := models.Subject{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&subject)
	res, err := ucon.srv.UpdateSubject(id, subject)
	if err == nil {
		return c.JSON(http.StatusCreated, response.SuccessResponse{
			Message: "Subject Updated",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
			Message: "Failed to Update Subject",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Delete Subject
// @Description Route Path for Delete Subject, for Administrator only.
// @Tags Subject
// @Accept json
// @Produce json
// @Param id path int true "subject id"
// @Success 200 {object} response.MessageOnlyResponse{} success
// @Failure 417 {object} response.ErrorResponse{} error
// @Router /subject/{id}/delete [DELETE]
func (ucon SubjectController) DeleteSubject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteSubject(id)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnlyResponse{
			Message: "Subject Updated",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.ErrorResponse{
			Message: "Failed to Update Subject",
			Error: err,
		})
	}
}

