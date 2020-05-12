package chatbot

import (
	"fmt"
	"os"

	"../model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// CreateNewBotConnection creates new connection to bot
// It passes token as argument and returns bot
func CreateNewBotConnection() (*tgbotapi.BotAPI, error) {
	os.Setenv("TOKEN_TG_BOT", "1101236908:AAGdRKCvt8EzpByAFjPKnof-gYKjdTE9jVM")
	if os.Getenv("TOKEN_TG_BOT") == "" {
		return nil, fmt.Errorf("invalid telegram token =(")
	}

	chatBot := model.Bot{
		Token: os.Getenv("TOKEN_TG_BOT"),
	}

	bot, err := tgbotapi.NewBotAPI(chatBot.Token)
	if err != nil {
		return nil, fmt.Errorf("CreateNewBotConnection error %w", err)
	}

	chatBot.Name = bot.Self.UserName

	return bot, nil
}
