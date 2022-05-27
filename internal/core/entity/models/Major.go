package models

type Major struct {
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	Code string `json:"code" example:"INF"`
	Name string `json:"name" example:"Informatics"`
}

func (*Major) TableName() string {
	return "majors"
}
