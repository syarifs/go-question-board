package database

import (
	"context"
	"go-question-board/internal/utils"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() (client *mongo.Client, isTest bool) {
	utils.LoadConfig()

	log.Println(os.Getenv("Env"))

	if os.Getenv("Env") == "testing" {
		client = nil
		isTest = true
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(utils.MONGODB_STRING))
	if err != nil {
		log.Fatal(err)
	}

	return
}
