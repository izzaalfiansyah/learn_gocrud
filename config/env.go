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
	if err == nil {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Println("Failed to read config")
		}
	} else {
		log.Println("No .env file found")
	}

	viper.SetDefault("APP_NAME", "GO CRUD LEARNING")
	viper.SetDefault("APP_PORT", 8000)

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Println("Error load .env config")
	}
}
