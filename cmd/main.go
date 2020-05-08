package main

import (
	"log"
	"os"

	"../chatbot"
	"../cmd/commands"
	"../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	if os.Getenv("TOKEN_TG_BOT") == "" {
		log.Println("Sorry! But invalid telegram token! =(")
		os.Exit(1)
	}
	chatBot := model.Bot{
		Token: os.Getenv("TOKEN_TG_BOT"),
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
