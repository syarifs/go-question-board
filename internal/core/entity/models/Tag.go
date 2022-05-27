package models

type Tag struct {
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	Name string `json:"name" example:"Year"`
	Value string `json:"value" example:"2019"`
}

func (*Tag) TableName() string {
	return "tags"
}
