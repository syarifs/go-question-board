package models

type (
	Subject struct {
		ID uint `json:"id" gorm:"primarykey"`
		Code string `json:"code"`
		Name string `json:"name"`
		MajorID int `json:"major_id"`
		Major Major
		Student []*User `json:"student" gorm:"many2many:student_subject"`
		Teacher []TeacherSubject `json:"teacher_class" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	}

	TeacherSubject struct {
		UserID *uint `json:"teacher_id" gorm:"primarykey;autoIncrement:false"`
		SubjectID uint `json:"subject_id" gorm:"primarykey;autoIncrement:false"`
		Class string `json:"class"`
		User User `json:"teacher"`
		Subject Subject `json:"subject"`
	}
)

func (*Subject) TableName() string {
	return "subjects"
}
