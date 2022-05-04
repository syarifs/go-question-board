package response

type SuccessResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type MessageOnlyResponse struct {
	Message string `json:"message"`
}
