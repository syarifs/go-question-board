package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	Name string `json:"name"`
	Value string `json:"value"`
}

func (*TagModel) TableName() string {
	return "tags"
}
