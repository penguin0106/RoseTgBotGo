package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config структура для хранения конфигураций бота.
type Config struct {
	TelegramBotToken string
	MongoDBURI       string
	// Добавьте другие конфигурационные параметры по необходимости
}

// LoadConfig загружает конфигурацию из файла .env.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	mongoDBURI := os.Getenv("MONGODB_URI")

	return &Config{
		TelegramBotToken: telegramBotToken,
		MongoDBURI:       mongoDBURI,
	}, nil
}
