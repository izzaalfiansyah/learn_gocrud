package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type env struct {
	AppName string `mapstructure:"APP_NAME"`
	AppPort int    `mapstructure:"APP_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

var Env env

func LoadConfig() {
	viper.AutomaticEnv()

	_, err := os.Stat(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	port := viper.GetInt("APP_PORT")

	if port == 0 {
		port = 8000
	}

	Env = env{
		AppName:    viper.GetString("APP_NAME"),
		AppPort:    port,
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
	}
}
