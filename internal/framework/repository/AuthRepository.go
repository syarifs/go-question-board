package repository

import (
	"context"
	"go-question-board/internal/core/entity/models"
	m "go-question-board/internal/core/entity/models"
	"go-question-board/internal/core/entity/response"
	"go-question-board/internal/utils/logger"

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

func (repo authRepository) Login(email string) (users response.UserDetails, err error) {

	var tags []models.Tag

	err = repo.sqldb.Model(models.User{}).
		Select("users.*, roles.name as role, majors.name as major").
		Joins("JOIN roles on roles.id = users.role_id").
		Joins("LEFT JOIN majors on majors.id = users.major_id").
		Where("email = ?", email).Scan(&users).Error

	if err != nil {
		return
	}

	err = repo.sqldb.Model(&models.User{}).
		Where("id = ?", users.ID).
		Scan(&users.Tags).Error
	
	users.Tags = tags
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
		{Key: "accesstoken", Value: token},
	}

	db := repo.mongo.Collection("token")
	_, err = db.DeleteOne(context.TODO(), filter)

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
			Value: bson.D{{Key: "refreshtoken", Value: new_token.RefreshToken}},
		},
		{
			Key: "$set",
			Value: bson.D{{Key: "accesstoken", Value: new_token.AccessToken}},
		},
	}


	db := repo.mongo.Collection("token")
	_, err = db.UpdateOne(context.TODO(), old_token, filter)

	if err != nil {
		logger.WriteLog(err)
	}

	return
}
