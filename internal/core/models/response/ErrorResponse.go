package response

type ErrorResponse struct {
	Error interface{} `json:"error"`
	Message string `json:"message"`
}
