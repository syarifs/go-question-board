package models

import "gorm.io/gorm"

type (
	Subject struct {
		ID uint `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		MajorID int `json:"major_id"`
		Major Major
		Teacher []TeacherSubject `json:"teacher_class" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	}

	TeacherSubject struct {
		ID uint `gorm:"primarykey"`
		UserID *int `json:"teacher_id"`
		SubjectID int `json:"subject_id"`
		Class string `json:"class"`
		User User `json:"teacher"`
		Subject Subject `json:"subject"`
	}
)

func (*Subject) BeforUpdate(db *gorm.DB) (err error) {
	return 
}

func (*Subject) TableName() string {
	return "subjects"
}
