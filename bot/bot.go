package bot

import (
	"RoseTgBotGo/database"
	"RoseTgBotGo/handlers"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	BotAPI      *tgbotapi.BotAPI
	DB          *database.MongoDB
	UpdatesChan chan<- tgbotapi.UpdatesChannel // Добавим поле UpdatesChan
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

	updates := make(chan tgbotapi.Update)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	go func() {
		updates, err := botAPI.GetUpdatesChan(u)
		if err != nil {
			log.Fatal(err)
		}
		for update := range updates {
			// Используйте канал UpdatesChan для отправки обновлений
			UpdatesChan <- update
		}
	}()

	return &Bot{
		BotAPI:      botAPI,
		DB:          db,
		UpdatesChan: updates,
	}, nil
}

func (b *Bot) Start() {
	for update := range b.UpdatesChan {
		if update.Message != nil {
			handlers.HandleMessage(b, update.Message)
		}
	}
}

func (b *Bot) SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.BotAPI.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
