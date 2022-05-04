package models

type Major struct {
	ID uint `json:"id" gorm:"primarykey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (*Major) TableName() string {
	return "majors"
}
