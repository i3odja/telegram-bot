package chatbot

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	telegramUpdateOffset  = 0
	telegramUpdateTimeout = 60
)

func GetUpdates(bot *tgbotapi.BotAPI) error {
	// channel initialization, that will be receive all updates from API
	updateConfig := tgbotapi.NewUpdate(telegramUpdateOffset)
	updateConfig.Timeout = telegramUpdateTimeout

	uch, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		return fmt.Errorf("getUpdates getting channel for updates error %v", err)
	}

	uch.Clear()

	for update := range uch {
		if update.CallbackQuery != nil {
			_, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))
			if err != nil {
				return fmt.Errorf("getUpdates AnswerCallbackQuery error %v", err)
			}

			m, err := bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
			if err != nil {
				return fmt.Errorf("getUpdates Send error %v", err)
			}

			update.Message = update.CallbackQuery.Message
			update.Message = &m
		}

		if update.Message != nil {
			err := SelectCommandsList(bot, &update)
			if err != nil {
				return fmt.Errorf("getUpdates %w", err)
			}
		}
	}

	return nil
}
