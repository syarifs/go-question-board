package response

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
	
	SubjectTeacher struct {
		Subject SubjectWithoutTeacher `json:"subject"`
		Class string `json:"class"`
	}
	
	SubjectWithoutTeacher struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
	}
	
	Subject struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Teacher []TeacherSubject `json:"teacher_class"`
	}
)
