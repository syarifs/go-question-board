package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name"`
	Value string `json:"value"`
}

func (*Tag) TableName() string {
	return "tags"
}
