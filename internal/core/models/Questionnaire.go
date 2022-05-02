package models

import (
	"gorm.io/gorm"
)

type Questionnaire struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []Tag `json:"tags" gorm:"many2many:questionnaire_tags"`
	Question []Question `json:"questions"`
	CreatedBy int `json:"created_by"`
	DonedBy int `json:"done_by"`
}

func (*Questionnaire) TableName() string {
	return "questionnaire"
}
