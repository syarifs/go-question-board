package models

import "time"

type Questionnaire struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []Tag `json:"tags" gorm:"many2many:questionnaire_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question []Question `json:"questions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedBy uint `json:"created_by"`
	Creator User `gorm:"foreignkey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Completor []User `gorm:"many2many:quest_user_complete;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (*Questionnaire) TableName() string {
	return "questionnaire"
}
