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

type UserController struct {
	srv *service.UserService
}

func NewUserController(srv *service.UserService) *UserController {
	return &UserController{srv}
}

// CreateResource godoc
// @Summary Create New User
// @Description Route Path for Insert New User, for Administrator only.
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  models.User{}  true "send request user code and user name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/user [post]
func (ucon UserController) CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := ucon.srv.CreateUser(user)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "User Created",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Create User",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Get All User
// @Description Route Path for Get List of User, for Administrator only.
// @Tags User
// @Security ApiKey
// @Success 200 {object} response.MessageData{data=[]response.User{}} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/user [get]
func (ucon UserController) ReadUser(c echo.Context) error {
	res, err := ucon.srv.ReadUser()
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "User Fetched",
			Data: res,
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Fetch User",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Get User By ID
// @Description Route Path for Get User Details By ID, for Administrator only.
// @Tags User
// @Security ApiKey
// @Param id path int true "user id"
// @Success 200 {object} response.MessageData{data=response.UserDetails} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/user/{id} [get]
func (ucon UserController) ReadUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ucon.srv.ReadUserByID(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "User Fetched",
			Data: user,
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Fetch User",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Update User
// @Description Route Path for Update User, for Administrator only.
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param body  body  models.User{}  true "send request user code and user name"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/user/{id}/update [PUT]
func (ucon UserController) UpdateUser(c echo.Context) error {
	user := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&user)
	err := ucon.srv.UpdateUser(id, user)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "User Updated",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Update User",
			Error: err.Error(),
		})
	}
}

// CreateResource godoc
// @Summary Delete User
// @Description Route Path for Delete User, for Administrator only.
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /admin/user/{id}/delete [DELETE]
func (ucon UserController) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteUser(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "User Deleted",
		})
	} else {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Delete User",
			Error: err.Error(),
		})
	}
}

