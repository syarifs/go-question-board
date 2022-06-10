package jwt

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/utils/config"
	"go-question-board/internal/utils/errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id float64, level string) (t models.Token, err error) {
	role := strings.ToLower(level)

	expTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["role"] = role
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t.AccessToken, err = token.SignedString(config.SERVER_SECRET)

	rexpTime := time.Now().Add(time.Hour * 24).Unix()
	rclaims := jwt.MapClaims{}
	rclaims["exp"] = rexpTime
	rclaims["iat"] = time.Now().Unix()
	rclaims["nbf"] = time.Now().Unix()
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)
	t.RefreshToken, err = rtoken.SignedString(config.SERVER_SECRET)
	return
}

func RefreshToken(token_string models.Token) (t models.Token, err error) {
	token, err := ExtractToken(token_string.RefreshToken)

	if _, ok := token.(jwt.MapClaims); ok {
		tkn, _ := ExtractToken(token_string.AccessToken)
		user_id := tkn.(jwt.MapClaims)["user_id"]
		role := tkn.(jwt.MapClaims)["role"]
		if user_id != nil && role != nil {
			return CreateToken(user_id.(float64), role.(string))
		}
	}
	return t, errors.New(500, "failed to generate new token")
}
