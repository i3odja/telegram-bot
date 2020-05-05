package main

import (
	"log"

	"../chatbot"
	"../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	botToken = "1161561075:AAG6WNCUAgAH0V-l5CG2QGo5smCzELERSow"
)

func main() {
	// setup chat bot token
	chatBot := model.Bot{
		Token: botToken,
	}

	// connection to bot with token
	bot, err := chatbot.CreateNewBotConnection(chatBot.Token)
	if err != nil {
		log.Println(err)
	}

	chatBot.Name = bot.Self.UserName

	// use bot.Debug equal to false to switch off debug mode
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// channel initialization, that will be receive all updates from API
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	uch, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatalf("getting channel for updates error %v", err)
	}
	uch.Clear()

	// reading updates from the channel (infinity loop)
	for {
		select {
		case update := <-uch:
			// setup userInfo structure with user information
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
		}

	}
}
