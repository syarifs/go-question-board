package repository

import (
	"go-question-board/internal/core/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, mongo *mongo.Database) *repository.Repository {
	return &repository.Repository{
		Auth: NewAuthRepository(db, mongo),
	}
}
