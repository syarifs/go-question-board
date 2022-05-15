package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
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

// CreateResource godoc
// @Summary Create New Major
// @Description Route Path for Insert New Major, for Administrator only.
// @Tags Major
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  models.Major{}  true "send request major code and major name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/major [post]
func (ucon MajorController) CreateMajor(c echo.Context) error {
	major := models.Major{}
	c.Bind(&major)
	err := ucon.srv.CreateMajor(major)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Major Created",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Create Major",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Get All Major
// @Description Route Path for Get List of Major, for Administrator only.
// @Tags Major
// @Security ApiKey
// @Success 200 {object} response.MessageData{data=[]models.Major} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/major [get]
func (ucon MajorController) ReadMajor(c echo.Context) error {
	res, err := ucon.srv.ReadMajor()
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Major Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch Major",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Update Major
// @Description Route Path for Update Major, for Administrator only.
// @Tags Major
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "major id"
// @Param body  body  models.Major{}  true "send request major code and major name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/major/{id}/update [PUT]
func (ucon MajorController) UpdateMajor(c echo.Context) error {
	major := models.Major{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&major)
	err := ucon.srv.UpdateMajor(id, major)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Major Updated",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Update Major",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Delete Major
// @Description Route Path for Delete Major, for Administrator only.
// @Tags Major
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "major id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/major/{id}/delete [DELETE]
func (ucon MajorController) DeleteMajor(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteMajor(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Major Deleted",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Delete Major",
			Error: err.Error(),
		})
	}
}

