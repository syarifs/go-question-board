package seeds

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func levelSeeder() Seed {
	seeds := []models.Role{
		{Name: "Administrator"},
		{Name: "Teacher"},
		{Name: "Student"},
	}
	model := &models.Role{}

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
