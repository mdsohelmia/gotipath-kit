package core

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	Environment   string `mapstructure:"ENV"`
	LogOutput     string `mapstructure:"LOG_OUTPUT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	DB_CONNECTION string `mapstructure:"mapstructure"`
	//Database connection
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_DATABASE string `mapstructure:"DB_DATABASE"`
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
}

// NewEnv creates a new environment
func NewEnv() Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return env
}
