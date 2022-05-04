package models

type Subject struct {
	ID uint `json:"id" gorm:"primarykey"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (*Subject) TableName() string {
	return "subjects"
}
