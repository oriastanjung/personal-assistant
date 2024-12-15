package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OPENAI_API_KEY string
	PORT           string
}

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config := &Config{
		OPENAI_API_KEY: os.Getenv("OPENAI_API_KEY"),
		PORT:           os.Getenv("PORT"),
	}
	return config
}
