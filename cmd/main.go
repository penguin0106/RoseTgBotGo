package cmd

import (
	"github.com/your_username/your_project_name/bot"
	"github.com/your_username/your_project_name/config"
	"log"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
		return
	}

	// Создание нового бота
	myBot, err := bot.NewBot(cfg.TelegramBotToken, cfg.MongoDBURI)
	if err != nil {
		log.Fatal("Error creating bot:", err)
		return
	}

	// Запуск бота
	myBot.Start()
}
