package commands

import (
	"fmt"
	"log"

	"../../chatbot"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func greeter(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	userInfo, err := chatbot.SetupUserInfo(update.Message.From)
	if err != nil {
		log.Println("SetupUserInfo error")
	}

	// Chat or dialog ID
	// It can be privet chat ID (then it is equal to UserID)
	// or public chat/channel ID
	ChatID := update.Message.Chat.ID

	// Message text
	Text := update.Message.Text

	log.Printf("[%s] %d %s", userInfo.Login, ChatID, Text)

	// Reply to user
	reply := chatbot.CreateReply(userInfo)

	// Create message
	msg := tgbotapi.NewMessage(ChatID, reply)
	// and send it
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("greeter Send error %w", err)
	}

	return nil
}
