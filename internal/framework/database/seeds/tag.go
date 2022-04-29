package seeds

import (
	"go-question-board/internal/core/models"

	"gorm.io/gorm"
)

func tagSeeder() Seed {
	seeds := []models.TagModel{
		{Name: "Year", Value: "2019"},
		{Name: "Year", Value: "2020"},
		{Name: "Year", Value: "2021"},
		{Name: "Class", Value: "A"},
		{Name: "Class", Value: "B"},
	}
	model := &models.TagModel{}

	return Seed{
		models: model,
		run: func(db *gorm.DB) (err error) {
			for _, v := range seeds {
				err = db.Create(&v).Error
			}
			return
		},
	}
}
