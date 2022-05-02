package models

import "gorm.io/gorm"

type Major struct {
	gorm.Model
	Code string `json:"code"`
	Name string `json:"name"`
}

func (*Major) TableName() string {
	return "majors"
}
