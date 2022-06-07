package repository

import (
	"context"
	m "go-question-board/internal/core/entity/models"
	"go-question-board/internal/utils/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authRepository struct {
	sqldb *gorm.DB
	mongo *mongo.Client
}

func NewAuthRepository(sqldb *gorm.DB, mongodb *mongo.Client) *authRepository {
	return &authRepository{sqldb: sqldb, mongo: mongodb}
}

func (repo authRepository) Login(email string) (users m.User, err error) {
	err = repo.sqldb.
		Preload(clause.Associations).
		Where("email = ?", email).First(&users).Error
	return
}

func (repo authRepository) SaveToken(token m.Token) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}

	db := repo.mongo.Database("question_board").Collection("token")
	_, err = db.InsertOne(context.TODO(), token)

	return
}

func (repo authRepository) RevokeToken(token m.Token) (err error) {
	// check and skip token saving if in testing mode
	if repo.mongo == nil {
		return
	}


	filter := bson.D{
		{Key: "accesstoken", Value: token.AccessToken},
	}

	db := repo.mongo.Database("question_board").Collection("token")
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


	db := repo.mongo.Database("question_board").Collection("token")
	_, err = db.UpdateOne(context.TODO(), old_token, filter)

	if err != nil {
		logger.WriteLog(err)
	}

	return
}
