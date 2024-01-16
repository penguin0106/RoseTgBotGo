package handlers

import (
	"RoseTgBotGo/bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

// HandleMessage метод для обработки входящего сообщения.
func HandleMessage(b *bot.Bot, message *tgbotapi.Message) {
	// Обработка входящего сообщения
	// Пример: регистрация нового пользователя
	if message.Text == "/register" {
		user, err := RegisterUser(b, message)
		if err != nil {
			log.Println("Error registering user:", err)
			return
		}
		b.SendMessage(b.BotAPI, message.Chat.ID, "Registration successful. Welcome, "+user.FirstName+"!")
	} else if message.Text == "/add_admin" {
		// Обработка команды добавления админа
		b.AdminPanel.HandleAddAdmin(message)
	}
	// Добавьте другие обработчики по необходимости
}

func RegisterUser(b *bot.Bot, message *tgbotapi.Message) (*database.User, error) {
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
