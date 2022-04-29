package seeds

import (
	"go-question-board/internal/core/models"

	"gorm.io/gorm"
)

func levelSeeder() Seed {
	seeds := []models.LevelModel{
		{Name: "Administrator"},
		{Name: "Teacher"},
		{Name: "Student"},
	}
	model := &models.LevelModel{}

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
