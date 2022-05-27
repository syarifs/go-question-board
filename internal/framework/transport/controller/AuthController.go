package controller

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/request"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/transport/middleware"
	"go-question-board/internal/utils/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)


type AuthController struct {
	srv *service.AuthService
}

func NewAuthController(srv *service.AuthService) *AuthController {
	return &AuthController{srv}
}

// CreateResource godoc
// @Summary Login
// @Description Login and get Authorization Token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param body  body  request.LoginRequest{}  true "send request email, password"
// @Success 200 {object} response.UserDetails{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /login [post]
func (acon AuthController) Login(c echo.Context) error {
	var login request.LoginRequest
	c.Bind(&login)
	res, err := acon.srv.Login(login)
	if err != nil {
		error := err.(*errors.RequestError)
		return c.JSON(error.Code(), response.Error{
			Message: "Failed to Log User In",
			Error: err.Error(),
		})
	}
	jwt, err := middleware.CreateToken(int(res.ID), res.Level.Name)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, response.Error{
			Message: "Failed to Create Authentication Token",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.MessageDataJWT{
		Message: "User Logged In",
		Data: res,
		JWT: jwt,
	})
}

// CreateResource godoc
// @Summary Refresh Token
// @Description Route Path for Get New Access Token
// @Tags Authorization
// @Accept json
// @Produce json
// @Param body  body  models.Token{}  true "send request access_token, refresh_token"
// @Success 200 {object} models.Token{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /refresh_token [post]
func (acon AuthController) RefreshToken(c echo.Context) error {
	rtoken := models.Token{}
	c.Bind(&rtoken)
	token, err := acon.srv.RefreshToken(rtoken)
	
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Refresh Token",
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Token Refreshed",
		"jwt": token,
	})
}

