package seeds

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func majorSeeder() Seed {
	seeds := []models.Major{
		{Code: "INF", Name: "Informatika"},
		{Code: "SI", Name: "Sistem Informasi"},
	}
	model := &models.Major{}

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
