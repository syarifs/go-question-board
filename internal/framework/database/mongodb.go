package database

import (
	"context"
	"go-question-board/internal/utils/config"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB() (db *mongo.Database) {

	if os.Getenv("Env") == "testing" {
		db = nil
		return
	}
	
	config.LoadConfig()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGODB_STRING))
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(config.MONGODB_DATABASE)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return
}
