package repository

import (
	"context"
	"errors"
	m "go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type authRepository struct {
	sqldb *gorm.DB
	mongo *mongo.Database
}

func NewAuthRepository(sqldb *gorm.DB, mongodb *mongo.Database) *authRepository {
	return &authRepository{sqldb: sqldb, mongo: mongodb}
}

func (repo authRepository) Login(email string) (users response.User, err error) {
	err = repo.sqldb.Model(&m.Profile{}).
		Select(`users.*, profiles.full_name, profiles.birth_date,
						profiles.gender, roles.name as roles`).
		Joins("left join users on users.id = profiles.user_id").
		Joins("left join roles on users.role_id = roles.id").
		Where("email = ?", email).Scan(&users).Error
	return
}

func (repo authRepository) SaveToken(token m.Token) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	db := repo.mongo.Collection("token")
	_, err = db.InsertOne(context.TODO(), token)

	return
}

func (repo authRepository) RevokeToken(token string) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	filter := bson.D{
		{Key: "access_token", Value: token},
	}

	db := repo.mongo.Collection("token")
	res := db.FindOneAndDelete(context.TODO(), filter)

	if res.Err() != nil {
		err = errors.New("invalid or expired token")
	}

	return
}

func (repo authRepository) UpdateToken(old_token m.Token, new_token m.Token) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	filter := bson.D{
		{
			Key: "$set",
			Value: bson.D{{Key: "refresh_token", Value: new_token.RefreshToken}},
		},
		{
			Key: "$set",
			Value: bson.D{{Key: "access_token", Value: new_token.AccessToken}},
		},
	}

	db := repo.mongo.Collection("token")
	res := db.FindOneAndUpdate(context.TODO(), old_token, filter)

	if res.Err() != nil {
		err = errors.New("invalid or expired token")
	}

	return
}
