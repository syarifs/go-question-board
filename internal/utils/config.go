package utils

import (
	"github.com/spf13/viper"
)

var (
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST string
	DB_DRIVER string
	SERVER_PORT string
	FIREBASE_ACCESS_KEY string
	SERVER_SECRET []byte
)

func LoadConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	DB_DATABASE = viper.Get("db.DATABASE").(string)
	DB_USERNAME = viper.Get("db.USERNAME").(string)
	DB_PASSWORD = viper.Get("db.PASSWORD").(string)
	DB_HOST = viper.Get("db.HOST").(string)
	DB_DRIVER = viper.Get("db.DRIVER").(string)
	SERVER_PORT = viper.Get("server.PORT").(string)
	FIREBASE_ACCESS_KEY = viper.Get("server.FIREBASE_ACCESS_KEY").(string)
	SERVER_SECRET= []byte(viper.GetString("server.SECRET"))
}
