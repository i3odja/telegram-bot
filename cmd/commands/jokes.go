package commands

import (
	"encoding/json"
	"fmt"
	"net/url"

	"../../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const jokesURL = "api.chucknorris.io/jokes/random"

func getJoke(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	urlJoke, err := url.Parse(jokesURL)
	if err != nil {
		return fmt.Errorf("getJoke Parse error %w", err)
	}

	jsonJoke, err := sendRequest(urlJoke)
	if err != nil {
		return fmt.Errorf("getJoke sendRequest error %w", err)
	}

	// Parse JSON
	dataJoke := new(model.Joke)
	err = json.Unmarshal(jsonJoke, &dataJoke)
	if err != nil {
		return fmt.Errorf("getJoke JSON Unmarshal error %w", err)
	}
	fmt.Println(dataJoke)

	text := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш анекдот готовий!")
	_, err = bot.Send(text)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	reply := makeReplyJoke(dataJoke)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	return nil
}

func makeReplyJoke(data *model.Joke) (reply string) {
	reply = data.Value

	return
}
