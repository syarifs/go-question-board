package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/service"
	"net/http"

	"strconv"
	"github.com/labstack/echo/v4"
)

type MajorController struct {
	srv *service.MajorService
}

func NewMajorController(srv *service.MajorService) *MajorController {
	return &MajorController{srv}
}

func (ucon MajorController) CreateMajor(c echo.Context) error {
	major := models.Major{}
	c.Bind(&major)
	res, err := ucon.srv.CreateMajor(major)
	if err == nil {
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "Major Created",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Create Major",
			"error": err.Error(),
		})
	}
}

func (ucon MajorController) ReadMajor(c echo.Context) error {
	res, err := ucon.srv.ReadMajor()
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Major Fetched",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Fetch Major",
			"error": err.Error(),
		})
	}
}

func (ucon MajorController) UpdateMajor(c echo.Context) error {
	major := models.Major{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&major)
	res, err := ucon.srv.UpdateMajor(id, major)
	if err == nil {
		return c.JSON(http.StatusAccepted, echo.Map{
			"message": "Major Updated",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Update Major",
			"error": err.Error(),
		})
	}
}

func (ucon MajorController) DeleteMajor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteMajor(id)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Major Deleted",
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Delete Major",
			"error": err.Error(),
		})
	}
}

