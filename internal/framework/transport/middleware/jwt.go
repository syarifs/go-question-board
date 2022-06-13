package middleware

import (
	"context"
	"go-question-board/internal/core/entity/response"
	ujwt "go-question-board/internal/utils/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database

func NewJWTConnection(mongo *mongo.Database) {
	client = mongo
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := ujwt.GetToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
		}

		filter := bson.D{
			{Key: "access_token", Value: token},
		}

		db := client.Collection("token")
		_, err = db.Find(context.TODO(), filter)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "invalid or expired token",
			})
		}

		return next(c)
	}
}

func AdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := ujwt.GetTokenData(c, "administrator")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
		}


		if role == nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for administrator",
			})
		}
		return next(c)
	}
}

func TeacherPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := ujwt.GetTokenData(c, "teacher")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
		}

		if role == nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: "access for this route only for teacher",
			})
		}
		return next(c)
	}
}

func StudentPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, err := ujwt.GetTokenData(c, "student")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.MessageOnly{
				Message: err.Error(),
			})
		}

		if role == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "access for this route only for student",
			})
		}
		return next(c)
	}
}
