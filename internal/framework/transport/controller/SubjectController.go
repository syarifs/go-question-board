package controller

import (
	"fmt"
	"go-question-board/internal/core/models"
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

func (ucon SubjectController) CreateSubject(c echo.Context) error {
	subject := models.Subject{}
	c.Bind(&subject)
	fmt.Println(subject)
	res, err := ucon.srv.CreateSubject(subject)
	if err == nil {
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "Subject Created",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Create Subject",
			"error": err,
		})
	}
}

func (ucon SubjectController) ReadSubject(c echo.Context) error {
	res, err := ucon.srv.ReadSubject()
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Subject Fetched",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Fetch Subject",
			"error": err,
		})
	}
}

func (ucon SubjectController) UpdateSubject(c echo.Context) error {
	subject := models.Subject{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&subject)
	fmt.Println(subject)
	res, err := ucon.srv.UpdateSubject(id, subject)
	if err == nil {
		return c.JSON(http.StatusAccepted, echo.Map{
			"message": "Subject Updated",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Update Subject",
			"error": err,
		})
	}
}

func (ucon SubjectController) DeleteSubject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteSubject(id)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Subject Deleted",
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Delete Subject",
			"error": err,
		})
	}
}

