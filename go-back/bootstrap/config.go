package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Env struct {
	// Common
	SERVER_PORT string
	SERVER_HOST string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_SSLMODE  string

	// Secrets
	SERVER_SECRET string
	DB_PASSWORD   string
}

func NewEnv() (*Env, error) {
	env := Env{}

	viper.SetConfigFile("config.yml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return &env, err
	}

	if err := godotenv.Load(); err != nil {
		return &env, err
	}

	env.SERVER_PORT = viper.GetString("SERVER.PORT")
	env.SERVER_HOST = viper.GetString("SERVER.HOST")
	env.DB_HOST = viper.GetString("DB.HOST")
	env.DB_PORT = viper.GetString("DB.PORT")
	env.DB_NAME = viper.GetString("DB.NAME")
	env.DB_USER = viper.GetString("DB.USER")
	env.DB_SSLMODE = viper.GetString("DB.SSLMODE")

	env.SERVER_SECRET = viper.GetString("SERVER.SECRET")
	env.DB_PASSWORD = viper.GetString("DB.PASSWORD")

	return &env, nil
}
