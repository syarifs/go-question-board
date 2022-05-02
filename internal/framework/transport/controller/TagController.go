package controller

import (
	"fmt"
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type TagController struct {
	srv *service.TagService
}

func NewTagController(srv *service.TagService) *TagController {
	return &TagController{srv}
}

func (ucon TagController) CreateTag(c echo.Context) error {
	tag := models.Tag{}
	c.Bind(&tag)
	fmt.Println(tag)
	res, err := ucon.srv.CreateTag(tag)
	if err == nil {
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "Tag Created",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Create Tag",
			"error": err,
		})
	}
}

func (ucon TagController) ReadTag(c echo.Context) error {
	res, err := ucon.srv.ReadTag()
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Tag Fetched",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Fetch Tag",
			"error": err,
		})
	}
}

func (ucon TagController) UpdateTag(c echo.Context) error {
	tag := models.Tag{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&tag)
	fmt.Println(tag)
	res, err := ucon.srv.UpdateTag(id, tag)
	if err == nil {
		return c.JSON(http.StatusAccepted, echo.Map{
			"message": "Tag Updated",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Update Tag",
			"error": err,
		})
	}
}

func (ucon TagController) DeleteTag(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteTag(id)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Tag Deleted",
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Delete Tag",
			"error": err,
		})
	}
}

