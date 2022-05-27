package seeds

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func subjectSeeder() Seed {
	seeds := []models.Subject{
		{Code: "BSPGM", Name: "Basic Programming", MajorID: 1},
		{Code: "PRALG", Name: "Programming Algorithm", MajorID: 1},
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
