package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"../../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func getWeather(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	update.Message.Text = "Thank you!"

	urlWeather, err := createURL()
	if err != nil {
		return fmt.Errorf("getWeather createURL error %w", err)
	}

	fmt.Println(urlWeather)

	jsonWeather, err := sendRequest(urlWeather)
	if err != nil {
		return fmt.Errorf("getWeather sendRequest error %w", err)
	}

	// Parse JSON
	dataWeather := new(model.DataWeather)
	err = json.Unmarshal(jsonWeather, &dataWeather)
	if err != nil {
		return fmt.Errorf("getWeather JSON Unmarshal error %w", err)
	}
	fmt.Println(dataWeather)

	reply := makeReplyWeather(dataWeather)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	return nil
}

func makeReplyWeather(data *model.DataWeather) string {
	reply := "Ваш прогноз погоди готовий!\n\n"

	reply += fmt.Sprintf("Місто \t%s", data.Name)

	y, m, d := time.Now().Date()
	reply += fmt.Sprintf("\nПрогноз на сьогодні %v-%v-%v", d, m, y)

	reply += fmt.Sprintf("\nТемпература повітря %d С", int(data.Main.Temperature))

	reply += fmt.Sprintf("\n%s", data.Weather[0].Description)

	reply += fmt.Sprintf("\nШвидкість вітру %v м/с", data.Wind.Speed)

	return reply
}

func createURL() (*url.URL, error) {
	weatherHost := "api.openweathermap.org/data/2.5/weather"
	appID := "db9a441fce153ac5701b2235510e4d1b"
	city := "Lviv"

	u, err := url.Parse(weatherHost)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url error: %w", err)
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url query error: %w", err)
	}

	q.Set("q", city)
	q.Set("APPID", appID)
	q.Set("units", "metric")

	u.RawQuery = q.Encode()

	return u, nil
}

func sendRequest(url *url.URL) ([]byte, error) {
	resp, err := http.Get("http://" + url.String())
	if err != nil {

		return nil, fmt.Errorf("sendRequest Get error %w", err)
	}
	defer resp.Body.Close()

	bs := make([]byte, 1014)

	for true {
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))
		bs = bs[:n]
		if n == 0 || err != nil {
			break
		}
	}

	return bs, nil
}
