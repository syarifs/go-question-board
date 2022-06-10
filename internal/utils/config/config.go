package config

import (
	"log"

	"github.com/integralist/go-findroot/find"
	"github.com/spf13/viper"
)

var (
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST string
	MONGODB_STRING string
	MONGODB_DATABASE string
	SERVER_PORT string
	SERVER_SECRET []byte
)

func LoadConfig() {
	root, err := find.Repo()
	if err != nil {
		log.Fatalf("fatal error = %v", err.Error())
	}

	viper.AddConfigPath(root.Path + "/config")
	viper.SetConfigName("config")
	
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	MONGODB_STRING = viper.GetString("mongo.STRING")
	MONGODB_DATABASE = viper.GetString("mongo.DATABASE")
	DB_DATABASE = viper.GetString("mysql.DATABASE")
	DB_USERNAME = viper.GetString("mysql.USERNAME")
	DB_PASSWORD = viper.GetString("mysql.PASSWORD")
	DB_HOST = viper.GetString("mysql.HOST")
	SERVER_PORT = viper.GetString("server.PORT")
	SERVER_SECRET= []byte(viper.GetString("server.SECRET"))
}
