package jwt

import (
	"errors"
	"fmt"
	"go-question-board/internal/utils/config"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetTokenData(c echo.Context, data string) (extracted interface{}, err error) {
	header := c.Request().Header.Get("Authorization")
	header = strings.Split(header, " ")[1]
	extract, err := ExtractToken(header)
	if extract != nil {
		extracted = extract.(jwt.MapClaims)[data]
	}
	return
}

func GetToken(c echo.Context) (header string, err error) {
	header = c.Request().Header.Get("Authorization")
	headers := strings.Split(header, " ")
	if header == "" || len(headers) < 2 {
		err = errors.New("No Token Provided")
		return
	}
	header = headers[1]
	_, err = ExtractToken(header)
	return
}

func ExtractToken(tkn string) (token interface{}, err error) {
	token, err = jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return config.SERVER_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return token.(*jwt.Token).Claims, nil
}
