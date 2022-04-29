package models

import "gorm.io/gorm"

type MajorModel struct {
	gorm.Model
	Code string `json:"code"`
	Name string `json:"name"`
}

func (*MajorModel) TableName() string {
	return "majors"
}
