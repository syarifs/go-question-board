package controller

import (
	"go-question-board/internal/core/models/request"
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
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  request.SubjectRequest{}  true "send request subject code and subject name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /subject [post]
func (ucon SubjectController) CreateSubject(c echo.Context) error {
	subject := request.SubjectRequest{}
	c.Bind(&subject)
	err := ucon.srv.CreateSubject(subject)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Subject Created",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Create Subject",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Get All Subject
// @Description Route Path for Get List of Subject, for Administrator only.
// @Tags Subject
// @Security ApiKey
// @Success 200 {object} response.MessageData{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /subject [get]
func (ucon SubjectController) ReadSubject(c echo.Context) error {
	res, err := ucon.srv.ReadSubject()
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Subject Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Subject",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Update Subject
// @Description Route Path for Update Subject, for Administrator only.
// @Tags Subject
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "subject id"
// @Param body  body  request.SubjectRequest{}  true "send request subject code and subject name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /subject/{id}/update [PUT]
func (ucon SubjectController) UpdateSubject(c echo.Context) error {
	subject := request.SubjectRequest{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&subject)
	err := ucon.srv.UpdateSubject(id, subject)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Subject Updated",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Update Subject",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Delete Subject
// @Description Route Path for Delete Subject, for Administrator only.
// @Tags Subject
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "subject id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /subject/{id}/delete [DELETE]
func (ucon SubjectController) DeleteSubject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteSubject(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Subject Deleted",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Delete Subject",
			Error: err.Error(),
		})
	}
}

