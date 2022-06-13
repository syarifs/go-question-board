package seeds

import (
	"go-question-board/internal/core/entity/models"
	"go-question-board/internal/utils"

	"gorm.io/gorm"
)

func userSeeder() Seed {
	password, _ := utils.HashPassword("password")
	seeds := []models.User{
		{
			Email: "admin@web.io",
			Password: password,
			RoleID: 1,
			Status: 1,
		},
		{
			Email: "teacher@web.io",
			Password: password,
			RoleID: 2,
			Status: 1,
		},
		{
			Email: "student@web.io",
			Password: password,
			RoleID: 3,
			Status: 1,
		},
	}
	model := &models.User{}

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
