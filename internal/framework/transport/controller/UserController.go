package controller

import (
	"go-question-board/internal/core/models"
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

func (ucon UserController) CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	res, err := ucon.srv.CreateUser(user)
	if err == nil {
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "User Created",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Create User",
			"error": err,
		})
	}
}

func (ucon UserController) ReadUser(c echo.Context) error {
	res, err := ucon.srv.ReadUser()
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User Fetched",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Fetch User",
			"error": err,
		})
	}
}

func (ucon UserController) ReadUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ucon.srv.ReadUserByID(id)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User Fetched",
			"data": user,
		})
	} else {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Fetch User",
			"error": err,
		})
	}
}

func (ucon UserController) UpdateUser(c echo.Context) error {
	user := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&user)
	res, err := ucon.srv.UpdateUser(id, user)
	if err == nil {
		return c.JSON(http.StatusAccepted, echo.Map{
			"message": "User Updated",
			"data": res,
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Update User",
			"error": err,
		})
	}
}

func (ucon UserController) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ucon.srv.DeleteUser(id)
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User Deleted",
		})
	} else {
		return c.JSON(http.StatusNotModified, echo.Map{
			"message": "Failed to Delete User",
			"error": err,
		})
	}
}

