package errors

import (
	"go-question-board/internal/core/entity/response"
	"reflect"
)

type HTTPCheck struct {
	Code int
	Result interface{}
}

func CheckHTTP(res interface{}, err error, msg string) (r HTTPCheck, ok bool) {
	r.Code = 200
	var error *RequestError
	var errMsg any

	if err != nil {
		if reflect.TypeOf(err).String() == "RequestError" {
			error = err.(*RequestError)
			r.Code = error.Code()
		}

		if reflect.ValueOf(err).Kind().String() == "map" {
			errMsg = err
		} else {
			errMsg = err.Error()
		}

		r.Result = response.Error{
			Message: "Failed to " + msg,
			Error: errMsg,
		}

		return
	}

	ok = true

	if res != nil {
		r.Result = response.MessageData{
			Message: "Success to " + msg,
			Data: res,
		}
		return
	}
	r.Result = response.MessageOnly{
			Message: "Success to " + msg,
	}

	return
}
