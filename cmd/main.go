package cmd

import (
	"RoseTgBotGo/bot"
	"RoseTgBotGo/config"
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

	defer myBot.DB.Close() // Закрытие соединения с базой данных при завершении работы

	// Запуск бота
	myBot.Start()
}
