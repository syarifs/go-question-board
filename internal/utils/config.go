package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/integralist/go-findroot/find"
	"github.com/spf13/viper"
)

var (
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST string
	MONGODB_STRING string
	SERVER_PORT string
	SERVER_SECRET []byte
)

func LoadConfig() {
	root, err := find.Repo()
	if err != nil {
		log.Fatalf("fatal error = %v", err.Error())
	}
	fmt.Printf("%+v", root)

	viper.AddConfigPath(root.Path + "/config")
	viper.SetConfigName("config")
	
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if os.Getenv("MONGODB_STRING") != "" {
		MONGODB_STRING = os.Getenv("MONGODB_STRING")
	} else {
		MONGODB_STRING = viper.GetString("mongo.STRING")
	}

	DB_DATABASE = viper.Get("mysql.DATABASE").(string)
	DB_USERNAME = viper.Get("mysql.USERNAME").(string)
	DB_PASSWORD = viper.Get("mysql.PASSWORD").(string)
	DB_HOST = viper.Get("mysql.HOST").(string)
	SERVER_PORT = viper.Get("server.PORT").(string)
	SERVER_SECRET= []byte(viper.GetString("server.SECRET"))
}
