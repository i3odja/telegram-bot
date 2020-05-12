package covid

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"../../../model"
	"../helper"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const covid19URL = "api.covid19api.com/dayone/country/ukraine/status/confirmed/live"

func Statistic(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	urlCovid, err := url.Parse(covid19URL)
	if err != nil {
		return fmt.Errorf("statistic Parse error %w", err)
	}

	jsonCovid, err := helper.SendRequest(urlCovid)
	if err != nil {
		return fmt.Errorf("statistic sendRequest error %w", err)
	}

	dataCovid := make([]model.Covid, 1)
	err = json.Unmarshal(jsonCovid, &dataCovid)
	if err != nil {
		return fmt.Errorf("covid JSON Unmarshal error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш звіт по Covid-19 готовий!"))
	if err != nil {
		return fmt.Errorf("statictic Send error %w", err)
	}

	y, m, d := time.Now().AddDate(0, 0, -1).Date()

	res := fmt.Sprintf("Станом на %v-%v-%v\n\n", d, m, y)
	for _, v := range MakeReplyCovid(dataCovid) {
		res += v
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, res))
	if err != nil {
		return fmt.Errorf("statistic Send error %w", err)
	}

	return nil
}

func MakeReplyCovid(data []model.Covid) (reply []string) {
	text := fmt.Sprintf("%v\n", data[len(data)-1].Country)
	text += fmt.Sprintf("Кількість випадків: %v\n", data[len(data)-1].Cases)
	text += fmt.Sprintf("Кількість випадків за добу: %v\n\n", data[len(data)-1].Cases-data[len(data)-2].Cases)
	reply = append(reply, text)

	return
}
