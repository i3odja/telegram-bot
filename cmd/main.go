package main

import (
	"log"

	"github.com/i3odja/telegram-bot/chatbot"
)

const debugMode = true

func main() {
	bot, err := chatbot.CreateNewBotConnection()
	if err != nil {
		log.Fatal("cannot connect to bot %w", err)
	}

	// use bot.Debug equal to false to switch off debug mode
	bot.Debug = debugMode
	log.Printf("Authorized on account %s", bot.Self.UserName)

	err = chatbot.GetUpdates(bot)
	if err != nil {
		log.Println(err)
	}
}
