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

	text := tgbotapi.NewMessage(update.Message.Chat.ID, "💵 Ваш курс валют на сьогодні готовий!")
	_, err = bot.Send(text)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	reply := MakeReplyCurrency(dataCurrency)

	res := fmt.Sprintf("Станом на %v\n\n", dataCurrency[0].Exchangedate)
	for _, v := range reply {
		res += v
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	return nil
}

func MakeReplyCurrency(data []model.Currency) (reply []string) {
	for _, v := range data {
		if v.CC == "USD" || v.CC == "EUR" || v.CC == "RUB" || v.CC == "PLN" {
			if v.CC == "USD" {
				v.TXT = fmt.Sprintf("🇺🇸 %v", v.TXT)
			}
			if v.CC == "EUR" {
				v.TXT = fmt.Sprintf("🇪🇺 %v", v.TXT)
			}
			if v.CC == "RUB" {
				v.TXT = fmt.Sprintf("🇷🇺 %v", v.TXT)
			}
			if v.CC == "PLN" {
				v.TXT = fmt.Sprintf("🇵🇱 %v", v.TXT)
			}
			text := fmt.Sprintf("%v\n", v.TXT)
			text += fmt.Sprintf("Курс: %v\n\n", v.Rate)
			reply = append(reply, text)
		}
	}

	return
}
