package currency

import (
	"encoding/json"
	"fmt"
	"net/url"

	"../../../model"
	"../helper"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	currencyURL = "bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json"

	currencyFormatUSD = "🇺🇸 %v"
	currencyFormatEUR = "🇪🇺 %v"
	currencyFormatRUB = "🇷🇺 %v"
	currencyFormatPLN = "🇵🇱 %v"
	currencyRate      = "Курс: %v\n\n"

	currencyCaption = "Станом на %v\n\n"

	currencyMessage = "💵 Ваш курс валют на сьогодні готовий!"
)

// Currency get actual currency USD, EUR, RUB and PLN
func Currency(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	urlCurrency, err := url.Parse(currencyURL)
	if err != nil {
		return fmt.Errorf("currency Parse error %w", err)
	}

	jsonCurrency, err := helper.SendRequest(urlCurrency)
	if err != nil {
		return fmt.Errorf("currency sendRequest error %w", err)
	}

	dataCurrency := make([]model.Currency, 1)

	err = json.Unmarshal(jsonCurrency, &dataCurrency)
	if err != nil {
		return fmt.Errorf("currency JSON Unmarshal error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, currencyMessage))
	if err != nil {
		return fmt.Errorf("currency Send error %w", err)
	}

	res := fmt.Sprintf(currencyCaption, dataCurrency[0].Exchangedate)
	for _, v := range MakeReplyCurrency(dataCurrency) {
		res += v
	}

	//for {
	//	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, res))
	//	if err != nil {
	//		continue
	//	}
	//	return nil
	//}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, res))
	if err != nil {
		return fmt.Errorf("currency Send error %w", err)
	}
	return nil
}

// MakeReplyCurrency creates response message with operation result that will be sent to user
func MakeReplyCurrency(data []model.Currency) (reply []string) {
	for _, v := range data {
		if v.CC == "USD" || v.CC == "EUR" || v.CC == "RUB" || v.CC == "PLN" {
			if v.CC == "USD" {
				v.TXT = fmt.Sprintf(currencyFormatUSD, v.TXT)
			}

			if v.CC == "EUR" {
				v.TXT = fmt.Sprintf(currencyFormatEUR, v.TXT)
			}

			if v.CC == "RUB" {
				v.TXT = fmt.Sprintf(currencyFormatRUB, v.TXT)
			}

			if v.CC == "PLN" {
				v.TXT = fmt.Sprintf(currencyFormatPLN, v.TXT)
			}

			text := fmt.Sprintf("%v\n", v.TXT)
			text += fmt.Sprintf(currencyRate, v.Rate)
			reply = append(reply, text)
		}
	}

	return
}
