package controller

import (
	"go-question-board/internal/core/models"
	"go-question-board/internal/core/models/request"
	"go-question-board/internal/core/service"
	"go-question-board/internal/framework/transport/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)


type AuthController struct {
	srv *service.AuthService
}

func NewAuthController(srv *service.AuthService) *AuthController {
	return &AuthController{srv}
}

func (acon AuthController) Login(c echo.Context) error {
	var login request.LoginRequest
	c.Bind(&login)
	res, err := acon.srv.Login(login)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Log User In",
			"error": err,
		})
	}
	jwt, err := middleware.CreateToken(res.Level.Name)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Log User In",
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User Logged In",
		"data": res,
		"jwt": jwt,
	})
}

func (acon AuthController) RefreshToken(c echo.Context) error {
	rtoken := models.TokenModel{}
	c.Bind(&rtoken)
	token, err := acon.srv.RefreshToken(rtoken)
	
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"message": "Failed to Refresh Token",
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Token Refreshed",
		"jwt": token,
	})
}

