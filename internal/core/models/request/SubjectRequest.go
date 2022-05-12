package request

import "go-question-board/internal/core/models"

type (
	Teacher struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Class string `json:"class"`
	}

	SubjectRequest struct {
		ID uint `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major models.Major
		Teacher []Teacher `json:"teacher"`
	}
) 
