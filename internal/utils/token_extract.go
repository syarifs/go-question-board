package utils

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetTokenData(c echo.Context, data string) interface{} {
	header := c.Request().Header.Get("Authorization")
	header = strings.Split(header, " ")[1]
	extract, err := ExtractToken(header)
	fmt.Println(err)
	return extract.(jwt.MapClaims)[data]
}

func ExtractToken(tkn string) (token interface{}, err error) {
	token, err = jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return SERVER_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	return token.(*jwt.Token).Claims, nil
}
