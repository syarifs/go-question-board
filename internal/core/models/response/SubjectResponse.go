package response

import m "go-question-board/internal/core/models"

type (
	Teacher struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Class string `json:"class"`
	}
	
	UserSubject struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Teacher Teacher `json:"teacher"`
	}

	Subject struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Teacher []Teacher `json:"teacher"`
	}
)
