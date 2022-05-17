package response

import "go-question-board/internal/core/models"

type (
	MessageData struct {
		Message string `json:"message"`
		Data interface{} `json:"data"`
	}

	MessageDataJWT struct {
		Message string `json:"message"`
		Data interface{} `json:"data"`
		JWT models.Token `json:"jwt"`
	}

	MessageOnly struct {
		Message string `json:"message"`
	}
) 
