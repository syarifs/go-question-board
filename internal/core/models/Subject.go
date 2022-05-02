package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Code string `json:"code"`
	Name string `json:"name"`
}

func (*Subject) TableName() string {
	return "subjects"
}
