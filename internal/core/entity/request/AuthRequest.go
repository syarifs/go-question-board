package request

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginRequest struct {
	Email string `json:"email" example:"admin@web.io"`
	Password string `json:"password" example:"admin"`
}

func (lr LoginRequest) Validate() error {
	return validation.ValidateStruct(&lr,
		validation.Field(&lr.Email, validation.Required, is.EmailFormat),
		validation.Field(&lr.Password, validation.Required, validation.RuneLength(3, 0)),
	)
}
