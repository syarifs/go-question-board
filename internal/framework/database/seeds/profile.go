package seeds

import (
	"go-question-board/internal/core/entity/models"

	"gorm.io/gorm"
)

func profileSeeder() Seed {
	seeds := []models.Profile{
		{
			FullName: "Administrator",
			Gender: "Male",
			BirthDate: "1994-02-12",
			UserID:    1,
		},
		{
			FullName: "Teacher A",
			Gender: "Male",
			BirthDate: "1995-04-05",
			UserID:    1,
		},
		{
			FullName: "Student A",
			Gender: "Male",
			BirthDate: "2001-07-02",
			UserID:    1,
		},

	}
	model := &models.Profile{}

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
