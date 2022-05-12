package response

type (
	MessageData struct {
		Message string `json:"message"`
		Data interface{} `json:"data"`
	}

	MessageOnly struct {
		Message string `json:"message"`
	}
) 
