package models

type Tag struct {
	ID uint `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Value string `json:"value"`
}

func (*Tag) TableName() string {
	return "tags"
}
