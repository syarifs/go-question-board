package middleware

import (
	"errors"
	"fmt"
	"go-question-board/internal/core/models"
	"go-question-board/internal/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: utils.SERVER_SECRET,
		TokenLookup: "header:x-auth-token",
	})
}

func CreateToken(level string) (t models.TokenModel, err error) {
	expTime := time.Now().Add(time.Minute * 1).Unix()
	claims := jwt.MapClaims{}
	claims["role"] = level
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t.AccessToken, err = token.SignedString(utils.SERVER_SECRET)

	expTime = time.Now().Add(time.Hour * 24).Unix()
	rclaims := jwt.MapClaims{}
	rclaims["exp"] = expTime
	rclaims["iat"] = time.Now().Unix()
	rclaims["nbf"] = time.Now().Unix()
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)
	t.RefreshToken, err = rtoken.SignedString(utils.SERVER_SECRET)
	return
}

func RefreshToken(token_string models.TokenModel) (t models.TokenModel, err error) {
	token, err := jwt.Parse(token_string.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return utils.SERVER_SECRET, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tkn, _ := jwt.Parse(token_string.AccessToken, nil)
		if role := tkn.Claims.(jwt.MapClaims)["role"]; role != nil {
			return CreateToken(role.(string))
		}
	}
	return t, errors.New("failed to generate new token")
}
