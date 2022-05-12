package models

type (
	Subject struct {
		ID uint `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		MajorID int `json:"major_id"`
		Major Major
		Teacher []TeacherSubject `json:"teacher"`
	}

	TeacherSubject struct {
		UserID int
		SubjectID int
		Class string `json:"class"`
		User User 
		Subject Subject
	}
)

func (*Subject) TableName() string {
	return "subjects"
}
