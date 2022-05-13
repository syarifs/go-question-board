package request

import m "go-question-board/internal/core/models"

type (
	Teacher struct {
		ID int `json:"id"`
		Name string `json:"name"`
	}
	
	TeacherSubject struct {
		User Teacher	`json:"teacher"`
		Class string `json:"class"`
	}
	
	SubjectRequest struct {
		ID int `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Teacher []TeacherSubject `json:"teacher_class"`
	}
) 
