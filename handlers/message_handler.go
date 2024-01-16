package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/penguin0106/RoseTgBotGo/bot"
	"log"
	"time"
)

func HandleMessage(b *bot.Bot, message *tgbotapi.Message) {
	// Обработка входящего сообщения
	// Пример: регистрация нового пользователя
	if message.Text == "/register" {
		user, err := registerUser(b, message)
		if err != nil {
			log.Println("Error registering user:", err)
			return
		}
		bot.SendMessage(b.BotAPI, message.Chat.ID, "Registration successful. Welcome, "+user.FirstName+"!")
	}
	// Добавьте другие обработчики по необходимости
}

func registerUser(b *bot.Bot, message *tgbotapi.Message) (*database.User, error) {
	// Логика регистрации пользователя в базе данных
	// Пример: сохранение пользователя в MongoDB
	user := &database.User{
		ID:           message.From.ID,
		FirstName:    message.From.FirstName,
		LastName:     message.From.LastName,
		PhoneNumber:  "your_phone_number",
		Role:         "user",
		Registration: time.Now(),
	}

	err := b.DB.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
