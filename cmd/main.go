package main

import (
	"log"

	"../chatbot"
	"../cmd/commands"
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
			err := commands.SelectCommandsList(bot, &update)
			if err != nil {
				log.Println(err)
			}
		}

	}
}
