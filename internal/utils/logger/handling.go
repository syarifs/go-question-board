package logger

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	mongoWriter struct {
    client *mongo.Database
	}
	
	logStruct struct {
		Timestamp int64 `json:"timestamp"`
		Logs string `json:"logs"`
	}
)

var LogDriver mongoWriter

func NewLogger(c *mongo.Database) {
	LogDriver = mongoWriter{client: c}
}

func (mw *mongoWriter) Write(p []byte) (n int, err error) {
	var db = mw.client.Collection("logs")
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
	if LogDriver.client != nil {
		log.SetOutput(&LogDriver)
		log.SetPrefix("")
	}

	log.Println(logs)
}

