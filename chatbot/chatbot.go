package chatbot

import (
	"fmt"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// CreateNewBotConnection creates new connection to bot
// It passes token as argument and returns bot
func CreateNewBotConnection() (*tgbotapi.BotAPI, error) {
	token, err := SetupToken()
	if err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func SetupToken() (string, error) {
	if os.Getenv("TOKEN_TG_BOT") == "" {
		return "", fmt.Errorf("wrong telegram token =(")
	}
	return os.Getenv("TOKEN_TG_BOT"), nil
}
