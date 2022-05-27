package middleware

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/utils"
	"go-question-board/internal/utils/errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: utils.SERVER_SECRET,
	})
}

func AdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := utils.GetTokenData(c, "role")

		if err := errors.CheckError(nil, err); err != nil {
			return err
		}

		if role != "Administrator" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "access for this route only for administrator",
			})
		}
		return next(c)
	}
}

func TeacherPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := utils.GetTokenData(c, "role")

		if err := errors.CheckError(nil, err); err != nil {
			return err
		}

		if role != "Teacher" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "access for this route only for teacher",
			})
		}
		return next(c)
	}
}

func StudentPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := utils.GetTokenData(c, "role")

		if err := errors.CheckError(nil, err); err != nil {
			return err
		}

		if role != "Student" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "access for this route only for student",
			})
		}
		return next(c)
	}
}

func CreateToken(id int, level string) (t models.Token, err error) {
	expTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["role"] = level
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t.AccessToken, err = token.SignedString(utils.SERVER_SECRET)

	rexpTime := time.Now().Add(time.Hour * 24).Unix()
	rclaims := jwt.MapClaims{}
	rclaims["exp"] = rexpTime
	rclaims["iat"] = time.Now().Unix()
	rclaims["nbf"] = time.Now().Unix()
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)
	t.RefreshToken, err = rtoken.SignedString(utils.SERVER_SECRET)
	return
}

func RefreshToken(token_string models.Token) (t models.Token, err error) {
	token, err := utils.ExtractToken(token_string.RefreshToken)
	if _, ok := token.(jwt.MapClaims); ok {
		tkn, _ := utils.ExtractToken(token_string.AccessToken)
		user_id := tkn.(jwt.MapClaims)["user_id"]
		role := tkn.(jwt.MapClaims)["role"]
		if user_id != nil && role != nil {
			return CreateToken(user_id.(int) ,role.(string))
		}
	}
	return t, errors.New(500, "failed to generate new token")
}


