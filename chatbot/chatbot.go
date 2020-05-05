package chatbot

import (
	"fmt"

	"../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	welcomeMessageUK = "Привіт %s! Як у тебе справи?"
	welcomeMessageEN = "Hello %s! How are you?"
)

// CreateNewBotConnection creates new connection to bot
// It passes token as argument and returns bot
func CreateNewBotConnection(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("CreateNewBotConnection error %w", err)
	}

	return bot, nil
}

// SetupUserInfo setups info about user
func SetupUserInfo(user *tgbotapi.User) (*model.User, error) {
	return &model.User{
		ID:           user.ID,
		Login:        user.UserName,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		LanguageCode: user.LanguageCode,
	}, nil
}

// CreatReply creates reply to user
// It checks what language is using user in order to create welcome message that can understand user
func CreateReply(userInfo *model.User) string {
	switch userInfo.LanguageCode {
	case "uk":
		return fmt.Sprintf(welcomeMessageUK, userInfo.FirstName)
	default:
		return fmt.Sprintf(welcomeMessageEN, userInfo.FirstName)
	}
}
