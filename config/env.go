package config

import (
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
		panic("Please create .env file first.")
	}

	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		panic("Error reading config file")
	}

	if err := viper.Unmarshal(&Env); err != nil {
		panic("Error unmarshalling config")
	}
}
