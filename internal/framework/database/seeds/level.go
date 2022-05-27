package seeds

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func levelSeeder() Seed {
	seeds := []models.Level{
		{Name: "Administrator"},
		{Name: "Teacher"},
		{Name: "Student"},
	}
	model := &models.Level{}

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
