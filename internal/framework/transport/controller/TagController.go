package controller

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/service"
	"go-question-board/internal/utils/errors"
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

// CreateResource godoc
// @Summary Create New Tag
// @Description Route Path for Insert New Tag, for Administrator only.
// @Tags Tag
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  models.Tag{}  true "send request tag code and tag name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/tag [post]
func (ucon TagController) CreateTag(c echo.Context) error {
	tag := models.Tag{}
	c.Bind(&tag)
	err := ucon.srv.CreateTag(tag)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "Tag Created",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Create Tag",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Get All Tag
// @Description Route Path for Get List of Tag, for Administrator only.
// @Tags Tag
// @Security ApiKey
// @Success 200 {object} response.MessageData{data=[]models.Tag{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/tag [get]
func (ucon TagController) ReadTag(c echo.Context) error {
	res, err := ucon.srv.ReadTag()
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "Tag Fetched",
			Data: res,
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Fetch Tag",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Update Tag
// @Description Route Path for Update Tag, for Administrator only.
// @Tags Tag
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "tag id"
// @Param body  body  models.Tag{}  true "send request tag code and tag name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/tag/{id}/update [PUT]
func (ucon TagController) UpdateTag(c echo.Context) error {
	tag := models.Tag{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&tag)
	err := ucon.srv.UpdateTag(id, tag)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Tag Updated",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Update Tag",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Delete Tag
// @Description Route Path for Delete Tag, for Administrator only.
// @Tags Tag
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "tag id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/tag/{id}/delete [DELETE]
func (ucon TagController) DeleteTag(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteTag(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "Tag Deleted",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Delete Tag",
			Error: err.Error(),
		})
	}
}

