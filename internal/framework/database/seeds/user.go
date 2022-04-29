package seeds

import (
	"go-question-board/internal/core/models"

	"gorm.io/gorm"
)

func userSeeder() Seed {
	major := 1
	seeds := []models.UserModel{
		{
			Name: "Administrator",
			Email: "admin@web.io",
			Password: "admin",
			LevelID: 1,
			Status: 1,
		},
		{
			Name: "Ach. Dafid",
			Email: "dafid@web.io",
			Password: "dosen",
			LevelID: 2,
			Status: 1,
			MajorID: &major,
		},
		{
			Name: "Syarif Ubaidillah",
			Email: "syarif@web.io",
			Password: "mahasiswa",
			LevelID: 3,
			Status: 1,
			MajorID: &major,
		},
	}
	model := &models.UserModel{}

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
