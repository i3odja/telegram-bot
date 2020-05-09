package commands

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"../../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const covid19URL = "api.covid19api.com/dayone/country/ukraine/status/confirmed/live"

func Covid19(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	urlCovid19, err := url.Parse(covid19URL)
	if err != nil {
		return fmt.Errorf("Covid19 Parse error %w", err)
	}

	jsonCovid19, err := sendRequest(urlCovid19)
	if err != nil {
		return fmt.Errorf("Covid19 sendRequest error %w", err)
	}

	// Parse JSON
	dataCovid19 := make([]model.Covid19, 1)
	err = json.Unmarshal(jsonCovid19, &dataCovid19)
	if err != nil {
		return fmt.Errorf("getJoke JSON Unmarshal error %w", err)
	}

	text := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш звіт по Covid-19 готовий!")
	_, err = bot.Send(text)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	reply := MakeReplyCovid19(dataCovid19)
	y, m, d := time.Now().AddDate(0, 0, -1).Date()
	res := fmt.Sprintf("Станом на %v-%v-%v\n\n", d, m, y)
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

func MakeReplyCovid19(data []model.Covid19) (reply []string) {
	text := fmt.Sprintf("%v\n", data[len(data)-1].Country)
	text += fmt.Sprintf("Кількість випадків: %v\n", data[len(data)-1].Cases)
	text += fmt.Sprintf("Кількість випадків за добу: %v\n\n", data[len(data)-1].Cases-data[len(data)-2].Cases)
	reply = append(reply, text)

	return
}
