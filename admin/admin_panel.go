package admin

import (
	"RoseTgBotGo/bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

// AdminPanel структура для представления админ панели.
type AdminPanel struct {
	Bot        *bot.Bot
	AdminsList []int64
}

// NewAdminPanel создает новый экземпляр AdminPanel.
func NewAdminPanel(b *bot.Bot) *AdminPanel {
	return &AdminPanel{
		Bot:        b,
		AdminsList: make([]int64, 0),
	}
}

// AddAdmin добавляет администратора в список.
func (ap *AdminPanel) AddAdmin(adminID int64) {
	// Проверка, чтобы избежать добавления одного и того же админа несколько раз
	for _, existingAdminID := range ap.AdminsList {
		if existingAdminID == adminID {
			return
		}
	}

	// Добавление админа в список
	ap.AdminsList = append(ap.AdminsList, adminID)
	log.Printf("Admin added: %d\n", adminID)
}

// HandleAddAdmin обрабатывает команду добавления админа.
func (ap *AdminPanel) HandleAddAdmin(message *tgbotapi.Message) {
	// Запрос ID пользователя для добавления в админы
	adminID := ap.AskForAdminID(message)
	if adminID != 0 {
		// Добавление в список админов
		ap.AddAdmin(adminID)
		// Используйте SendMessage из вашего пакета bot
		ap.Bot.SendMessage(message.Chat.ID, "Admin added successfully.")
	} else {
		// Используйте SendMessage из вашего пакета bot
		ap.Bot.SendMessage(message.Chat.ID, "Admin not added.")
	}
}

// AskForAdminID запрашивает ID пользователя для добавления в админы.
func (ap *AdminPanel) AskForAdminID(message *tgbotapi.Message) int64 {
	// Отправить сообщение с вспомогательной клавиатурой
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите ID пользователя для добавления в админы:")
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Отмена"),
		),
	)
	msg.ReplyMarkup = keyboard

	// Отправить сообщение
	sentMsg, err := ap.Bot.BotAPI.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
		return 0
	}

	// Ожидать ответа от пользователя
	update := <-ap.UpdatesChan
	if update.Message.Text == "Отмена" {
		// Обработать отмену (не добавлять пользователя в админы)
		return 0
	}

	// Парсинг ID пользователя из ответа
	adminID, err := strconv.ParseInt(update.Message.Text, 10, 64)
	if err != nil {
		// Обработать ошибку парсинга
		log.Println("Error parsing admin ID:", err)
		return 0
	}

	// Добавление админа в список
	ap.AddAdmin(adminID)
	// Используйте SendMessage из вашего пакета bot
	ap.Bot.SendMessage(sentMsg.Chat.ID, "Admin added successfully.")
	return adminID
}
