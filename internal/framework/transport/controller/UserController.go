package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/response"
	"go-question-board/internal/core/service"
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
// @Router /user [post]
func (ucon UserController) CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := ucon.srv.CreateUser(user)
	if err == nil {
		return c.JSON(http.StatusCreated, response.MessageOnly{
			Message: "User Created",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Create User",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get All User
// @Description Route Path for Get List of User, for Administrator only.
// @Tags User
// @Security ApiKey
// @Success 200 {object} response.MessageData{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /user [get]
func (ucon UserController) ReadUser(c echo.Context) error {
	res, err := ucon.srv.ReadUser()
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "User Fetched",
			Data: res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch User",
			Error: err,
		})
	}
}

// CreateResource godoc
// @Summary Get User By ID
// @Description Route Path for Get List of User, for Administrator only.
// @Tags User
// @Security ApiKey
// @Success 200 {object} response.MessageData{} success
// @Failure 417 {object} response.Error{} error
// @Failure 400 {object} response.MessageOnly{} error
// @Failure 401 {object} response.MessageOnly{} error
// @Router /user/{id} [get]
func (ucon UserController) ReadUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ucon.srv.ReadUserByID(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageData{
			Message: "User Fetched",
			Data: user,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Fetch User",
			Error: err,
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
// @Router /user/{id}/update [PUT]
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
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Update User",
			Error: err,
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
// @Router /user/{id}/delete [DELETE]
func (ucon UserController) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteUser(id)
	if err == nil {
		return c.JSON(http.StatusOK, response.MessageOnly{
			Message: "User Deleted",
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Delete User",
			Error: err,
		})
	}
}

