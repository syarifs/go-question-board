package models

import "gorm.io/gorm"

type SubjectModel struct {
	gorm.Model
	Code string `json:"code"`
	Name string `json:"name"`
}

func (*SubjectModel) TableName() string {
	return "subjects"
}
