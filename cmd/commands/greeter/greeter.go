package greeter

import (
	"fmt"
	"log"

	"../../../model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	welcomeMessageUK = "Привіт %s! Як у тебе справи?"
	welcomeMessageEN = "Hello %s! How are you?"
)

func Greeter(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	userInfo, err := setupUserInfo(update.Message.From)
	if err != nil {
		log.Println("SetupUserInfo error")
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, CreateReply(userInfo)))
	if err != nil {
		return fmt.Errorf("greeter Send error %w", err)
	}

	return nil
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

// setupUserInfo setups info about user
func setupUserInfo(user *tgbotapi.User) (*model.User, error) {
	return &model.User{
		ID:           user.ID,
		Login:        user.UserName,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		LanguageCode: user.LanguageCode,
	}, nil
}
