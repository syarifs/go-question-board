package request

import (
	"go-question-board/internal/core/entity/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	BirthDate string `json:"birthdate"`
	Gender string `json:"gender"`
	Role models.Role `json:"role"`
	Status int `json:"status"`
	IsEdit bool `json:"-"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required.When(u.IsEdit), is.Email),
		validation.Field(&u.Name, validation.Required, validation.RuneLength(3, 0)),
		validation.Field(&u.BirthDate, validation.Required, validation.Date("2006-06-02")),
		validation.Field(&u.Gender, validation.Required),
		validation.Field(&u.Role, validation.Required),
	)
}

