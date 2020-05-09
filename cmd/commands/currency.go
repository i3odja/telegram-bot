package commands

import (
	"encoding/json"
	"fmt"
	"net/url"

	"../../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const currencyURL = "bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json"

func Currency(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	urlCurrency, err := url.Parse(currencyURL)
	if err != nil {
		return fmt.Errorf("getJoke Parse error %w", err)
	}

	jsonCurrency, err := sendRequest(urlCurrency)
	if err != nil {
		return fmt.Errorf("getJoke sendRequest error %w", err)
	}

	// Parse JSON
	dataCurrency := make([]model.Currency, 1)
	err = json.Unmarshal(jsonCurrency, &dataCurrency)
	if err != nil {
		return fmt.Errorf("getJoke JSON Unmarshal error %w", err)
	}

	text := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш курс валют на сьогодні готовий!")
	_, err = bot.Send(text)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	reply := MakeReplyCurrency(dataCurrency)

	for _, v := range reply {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, v)
		_, err = bot.Send(msg)
		if err != nil {
			return fmt.Errorf("getWeather Send error %w", err)
		}
	}

	return nil
}

func MakeReplyCurrency(data []model.Currency) (reply []string) {
	for _, v := range data {
		if v.CC == "USD" {
			text := fmt.Sprintf("%v\n", v.TXT)
			text += fmt.Sprintf("Станом на %v\n", v.Exchangedate)
			text += fmt.Sprintf("Курс: %v\n", v.Rate)
			reply = append(reply, text)
		}
		if v.CC == "EUR" {
			text := fmt.Sprintf("%v\n", v.TXT)
			text += fmt.Sprintf("Станом на %v\n", v.Exchangedate)
			text += fmt.Sprintf("Курс: %v\n", v.Rate)
			reply = append(reply, text)
		}
	}

	return
}
