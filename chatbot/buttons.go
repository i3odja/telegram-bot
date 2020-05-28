package chatbot

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func KeyboardButtons(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	keys := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Показати всі команди ❓", "/help"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Привітатися зі мною 👋", "/hello"),
			tgbotapi.NewInlineKeyboardButtonData("Звіт по корона-вірусі 裂", "/covid"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Прогноз погоди 🌤", "/weather"),
			tgbotapi.NewInlineKeyboardButtonData("Курс валют 💰", "/currency"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Жарти про Чак Норіса 😂", "/joke"),
			tgbotapi.NewInlineKeyboardButtonData("Показати зображення 🌄", "/picture"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Будь ласка, оберіть потрібну дію:")
	msg.ReplyMarkup = keys

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("buttons Send error %w", err)
	}

	return nil
}
