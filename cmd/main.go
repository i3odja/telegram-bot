package main

import (
	"log"

	"../chatbot"
)

const (
	debugMode = true
)

func main() {
	bot, err := chatbot.CreateNewBotConnection()
	if err != nil {
		log.Fatal(err)
	}

	// use bot.Debug equal to false to switch off debug mode
	bot.Debug = debugMode
	log.Printf("Authorized on account %s", bot.Self.UserName)

	err = chatbot.GetUpdates(bot)
	if err != nil {
		log.Println(err)
	}
}
