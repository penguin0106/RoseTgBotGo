package bot

import (
	"RoseTgBotGo/database"
	"RoseTgBotGo/handlers"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	BotAPI *tgbotapi.BotAPI
	DB     *database.MongoDB
	// Добавьте другие поля по необходимости
}

func NewBot(token, mongoURI string) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	db, err := database.NewMongoDB(mongoURI)
	if err != nil {
		return nil, err
	}

	return &Bot{
		BotAPI: botAPI,
		DB:     db,
	}, nil
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.BotAPI.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message != nil {
			handlers.HandleMessage(b, update.Message)
		}
	}
}
