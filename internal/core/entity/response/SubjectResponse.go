package response

import m "go-question-board/internal/core/entity/models"

type (
	Teacher struct {
		ID int `json:"id"`
		Name string `json:"name"`
	}
	
	Student struct {
		ID        uint `gorm:"primary_key"`
		Email string `json:"email"`
		Name string `json:"name"`
		Tags []m.Tag `json:"tags"`
		Major m.Major `json:"major"`
	}

	TeacherClass struct {
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
	
	SubjectWithTeacher struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Teacher []TeacherClass `json:"teacher_class"`
	}
	
	SubjectWithStudent struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Student []Student `json:"student"`
	}
	
	Subject struct {
		ID int `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		Major m.Major `json:"major"`
		Teacher []TeacherClass `json:"teacher_class"`
		Student []Student `json:"student"`
	}
)
