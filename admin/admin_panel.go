package admin

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/penguin0106/RoseTgBotGo/bot"
	"log"
)

// AdminPanel структура для представления админ панели.
type AdminPanel struct {
	Bot *bot.Bot
}

func NewAdminPanel(b *bot.Bot) *AdminPanel {
	return &AdminPanel{
		Bot: b,
	}
}

// HandleCommand обработчик команд от администратора.
func (ap *AdminPanel) HandleCommand(command string, message *tgbotapi.Message) {
	switch command {
	case "/add_admin":
		ap.handleAddAdmin(message)
	case "/send_broadcast":
		ap.handleSendBroadcast(message)
	default:
		// Добавьте обработку других команд администратора по необходимости
	}
}

func (ap *AdminPanel) handleAddAdmin(message *tgbotapi.Message) {
	// Реализуйте логику добавления администратора
	// Пример: запрос ID пользователя и добавление в список админов
	adminID := ap.askForAdminID(message)
	if adminID != 0 {
		// Добавление в список админов
		ap.Bot.AdminPanel.AddAdmin(adminID)
		ap.Bot.SendMessage(ap.BotAPI, message.Chat.ID, "Admin added successfully.")
	} else {
		ap.Bot.SendMessage(ap.BotAPI, message.Chat.ID, "Invalid user ID. Admin not added.")
	}
}

func (ap *AdminPanel) handleSendBroadcast(message *tgbotapi.Message) {
	// Реализуйте логику отправки рассылки
	// Пример: запрос текста рассылки и отправка всем пользователям
	text := ap.askForBroadcastText(message)
	if text != "" {
		users, err := ap.Bot.DB.GetAllUsers()
		if err != nil {
			log.Println("Error getting users for broadcast:", err)
			return
		}

		for _, user := range users {
			ap.Bot.SendMessage(ap.BotAPI, user.ID, text)
		}

		ap.Bot.SendMessage(ap.BotAPI, message.Chat.ID, "Broadcast sent successfully.")
	} else {
		ap.Bot.SendMessage(ap.BotAPI, message.Chat.ID, "Broadcast text cannot be empty.")
	}
}

func (ap *AdminPanel) askForAdminID(message *tgbotapi.Message) int64 {
	// Реализуйте логику запроса ID пользователя для добавления в администраторы
	// Верните 0, если ID неверен
	// Пример: отобразите клавиатуру для ввода ID и верните его значение
	return 0
}

func (ap *AdminPanel) askForBroadcastText(message *tgbotapi.Message) string {
	// Реализуйте логику запроса текста рассылки
	// Верните пустую строку, если текст неверен
	// Пример: отобразите клавиатуру для ввода текста и верните его значение
	return ""
}
