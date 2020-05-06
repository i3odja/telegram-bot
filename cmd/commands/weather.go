package commands

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func getWeather(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	update.Message.Text = "Thank you!"

	reply := "Дякую Вам за звернення! Прогноз погоди ще на стадії розробки..."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

	_, err := bot.Send(msg)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	return nil
}
