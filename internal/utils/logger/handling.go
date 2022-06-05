package logger

import (
	"context"
	"go-question-board/internal/framework/database"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	mongoWriter struct {
    client *mongo.Client
	}
	
	logStruct struct {
		Timestamp int64 `json:"timestamp"`
		Logs string `json:"logs"`
	}
)

var client, isTest = database.InitMongoDB() 

var LogDriver = mongoWriter{client: client}
var Logger = log.New(&LogDriver, "", 0)

func (mw *mongoWriter) Write(p []byte) (n int, err error) {
	var db = mw.client.Database("hospital").Collection("logs")
	doc := logStruct{
		Timestamp: time.Now().Unix(),
		Logs: string(p),
	}

	_, err = db.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	return len(p), nil
}

func WriteLog(logs interface{}) {
	if !isTest {
		Logger.Println(logs)
	} else {
		log.Println(logs)
	}
}

