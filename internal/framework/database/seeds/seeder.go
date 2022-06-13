package seeds

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Seed struct {
	models interface{}
	run func(*gorm.DB) error
}

func NewSeeders(db *gorm.DB) {
	for _, v := range seedAll() {
		if db.Migrator().HasTable(&v.models){
			err := db.First(&v.models).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := v.run(db); err != nil {
					log.Fatalf("Running seed error %e", err)
				}
			}
		}
	}
}

func seedAll() []Seed {
	return []Seed{
		levelSeeder(),
		userSeeder(),
		profileSeeder(),
	}
}

