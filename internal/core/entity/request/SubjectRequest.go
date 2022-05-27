package request

type (
	TeacherSubject struct {
		User User	`json:"teacher"`
		Class string `json:"class" example:"A"`
	}
	
	SubjectRequest struct {
		ID uint `json:"id" example:"1"`
		Code string `json:"code" example:"INF2213"`
		Name string `json:"name" example:"Logic Informatic"`
		MajorID int `json:"major_id"`
		Teacher []TeacherSubject `json:"teacher_class"`
		Student []User `json:"student"`
	}
) 
