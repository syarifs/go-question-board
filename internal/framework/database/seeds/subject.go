package seeds

import (
	"go-question-board/internal/core/models"

	"gorm.io/gorm"
)

func subjectSeeder() Seed {
	seeds := []models.Subject{
		{Code: "BSPGM", Name: "Basic Programming"},
		{Code: "PRALG", Name: "Programming Algorithm"},
	}
	model := &models.Subject{}

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
