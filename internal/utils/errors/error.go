package errors


func New(code int, text string) error {
	return &RequestError{
		code: code,
		message: text,
	}
}

type RequestError struct {
	code int
	message string
}

func (e *RequestError) Error() string {
	return e.message
}

func (e *RequestError) Code() int {
	return e.code
}
