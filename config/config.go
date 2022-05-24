package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port          string
	From          string
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_USERNAME string
	SMTP_PASSWORD string
}

func InitConfig() Config {
	var config Config
	godotenv.Load()

	config.Port = os.Getenv("PORT")
	config.SMTP_HOST = os.Getenv("SMTP_HOST")
	config.SMTP_PORT = os.Getenv("SMTP_PORT")
	config.SMTP_USERNAME = os.Getenv("SMTP_USERNAME")
	config.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	config.From = os.Getenv("FROMEMAIL")

	return config

}
