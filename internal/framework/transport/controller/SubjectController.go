package controller

import (
	"go-question-board/internal/core/models"
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
// @Router /admin/subject [post]
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
// @Success 200 {object} response.MessageData{data=[]response.SubjectWithoutTeacher{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/subject [get]
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
// @Summary Get Subject By ID
// @Description Route Path for Get Subject Details By ID.
// @Tags Subject
// @Security ApiKey
// @Param id  path  int  true "subject id"
// @Success 200 {object} response.MessageData{data=[]response.Subject{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/subject/{id} [get]
func (ucon SubjectController) ReadSubjectByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := ucon.srv.ReadSubjectByID(id)
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
// @Router /admin/subject/{id}/update [PUT]
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
// @Router /admin/subject/{id}/delete [DELETE]
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

// CreateResource godoc
// @Summary Get Teacher Subject
// @Description Route Path for Get List of Teacher Subject, for Teacher only.
// @Tags Subject
// @Security ApiKey
// @Param body body models.User{} true "user data for fetch subject"
// @Success 200 {object} response.MessageData{data=[]response.SubjectWithStudent{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /teacher/subject [get]
func (ucon SubjectController) ReadTeacherSubject(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	res, err := ucon.srv.ReadTeacherSubject(int(user.ID))
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
// @Summary Get Student Subject
// @Description Route Path for Get List of Student Subject, for Student only.
// @Tags Subject
// @Security ApiKey
// @Param body body models.User{} true "user data for fetch subject"
// @Success 200 {object} response.MessageData{data=[]response.SubjectWithTeacher{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /student/subject [get]
func (ucon SubjectController) ReadStudentSubject(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	res, err := ucon.srv.ReadStudentSubject(int(user.ID))
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
