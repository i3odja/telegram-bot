package currency

import (
	"encoding/json"
	"fmt"
	"net/url"
	
	"github.com/i3odja/telegram-bot/cmd/commands/helper"
	"github.com/i3odja/telegram-bot/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	currencyURL = "bank.gov.ua/NBUStatService/v1/statdirectory/exchange"

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
	values := []string{"EUR", "USD", "RUB", "PLN"}

	res := ""
	jsonCurrency := make([][]byte, 0)

	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, currencyMessage))
	if err != nil {
		return fmt.Errorf("currency Send error %w", err)
	}
	for i, v := range values {
		u, err := createCurrencyURL(v)
		if err != nil {
			return fmt.Errorf("jsonParse createURL error %w", err)
		}

		jsonCurrencyItem, err := helper.SendRequest(u)
		if err != nil {
			return fmt.Errorf("currency sendRequest error %w", err)
		}

		jsonCurrency = append(jsonCurrency, jsonCurrencyItem)

		dataCurrency := make([]model.Currency, len(values))

		err = json.Unmarshal(jsonCurrency[i], &dataCurrency)
		fmt.Println(&dataCurrency)
		if err != nil {
			return fmt.Errorf("currency JSON Unmarshal error %w", err)
		}

		for _, v := range MakeReplyCurrency(dataCurrency) {
			res += v
		}
	}

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

func createCurrencyURL(value string) (*url.URL, error) {
	u, err := url.Parse(currencyURL)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url error: %w", err)
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url query error: %w", err)
	}

	q.Set("valcode", value)
	q.Set("json", "")

	u.RawQuery = q.Encode()

	return u, nil
}
