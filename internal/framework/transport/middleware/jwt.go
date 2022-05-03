package middleware

import (
	"errors"
	"fmt"
	"go-question-board/internal/core/models"
	"go-question-board/internal/utils"
	"net/http"
	"strings"
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
		header := c.Request().Header.Get("Authorization")
		header = strings.Split(header, " ")[1]
		extract, _ := extractToken(header)
		if extract.(jwt.MapClaims)["role"] != "Administrator" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "access blocked. for administartor only",
			})
		}
		return next(c)
	}
}

func CreateToken(level string) (t models.Token, err error) {
	expTime := time.Now().Add(time.Minute * 15).Unix()
	claims := jwt.MapClaims{}
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

func extractToken(tkn string) (token interface{}, err error) {
	token, err = jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return utils.SERVER_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return token.(*jwt.Token).Claims, nil
}

func RefreshToken(token_string models.Token) (t models.Token, err error) {
	token, err := extractToken(token_string.RefreshToken)
	if _, ok := token.(jwt.MapClaims); ok {
		tkn, _ := extractToken(token_string.AccessToken)
		if role := tkn.(jwt.MapClaims)["role"]; role != nil {
			return CreateToken(role.(string))
		}
	}
	return t, errors.New("failed to generate new token")
}


