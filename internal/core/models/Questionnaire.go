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
	CreatedBy uint `json:"created_by"`
	CompletedBy *uint `json:"completed_by"`
	Creator User `gorm:"foreignkey:CreatedBy;references:ID"`
	Completor User `gorm:"foreignkey:CompletedBy;references:ID"`
}

func (*Questionnaire) TableName() string {
	return "questionnaire"
}
