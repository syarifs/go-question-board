package models

import "time"

type Questionnaire struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []Tag `json:"tags" gorm:"many2many:questionnaire_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question []Question `json:"questions"`
	CreatedBy uint `json:"created_by"`
	CompletedBy *uint `json:"completed_by"`
	Creator User `gorm:"foreignkey:CreatedBy;references:ID"`
	Completor []User `gorm:"many2many:quest_user_complete;foreignkey:CompletedBy;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (*Questionnaire) TableName() string {
	return "questionnaire"
}
