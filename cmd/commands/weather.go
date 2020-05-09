package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"../../model"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const weatherHost = "api.openweathermap.org/data/2.5/weather"

func getWeather(city string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	update.Message.Text = "Thank you!"

	urlWeather, err := createURL(city)
	if err != nil {
		return fmt.Errorf("getWeather createURL error %w", err)
	}

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

	text := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш прогноз погоди готовий!\n\n")
	_, err = bot.Send(text)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	reply := makeReplyWeather(dataWeather)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("getWeather Send error %w", err)
	}

	return nil
}

func makeReplyWeather(data *model.DataWeather) (reply string) {
	temp, minTemp, maxTemp, feelsTemp := "", "", "", ""

	reply += fmt.Sprintf("Місто \t%s", data.Name)

	y, m, d := time.Now().Date()
	reply += fmt.Sprintf("\nПрогноз погоди на сьогодні %v-%v-%v", d, m, y)

	switch {
	case int(data.Main.Temperature) > 0:
		temp = fmt.Sprintf("+%v", int(data.Main.Temperature))
	case int(data.Main.Temperature) < 0:
		temp = fmt.Sprintf("-%v", int(data.Main.Temperature))
	default:
		temp = fmt.Sprintf("%v", int(data.Main.Temperature))
	}

	switch {
	case int(data.Main.TemperatureMin) > 0:
		minTemp = fmt.Sprintf("+%v", int(data.Main.TemperatureMin))
	case int(data.Main.TemperatureMin) < 0:
		minTemp = fmt.Sprintf("-%v", int(data.Main.TemperatureMin))
	default:
		minTemp = fmt.Sprintf("%v", int(data.Main.TemperatureMin))
	}

	switch {
	case int(data.Main.TemperatureMax) > 0:
		maxTemp = fmt.Sprintf("+%v", int(data.Main.TemperatureMax))
	case int(data.Main.TemperatureMin) < 0:
		maxTemp = fmt.Sprintf("-%v", int(data.Main.TemperatureMax))
	default:
		maxTemp = fmt.Sprintf("%v", int(data.Main.TemperatureMax))
	}

	switch {
	case int(data.Main.TemperatureFeels) > 0:
		feelsTemp = fmt.Sprintf("+%v", int(data.Main.TemperatureFeels))
	case int(data.Main.TemperatureFeels) < 0:
		feelsTemp = fmt.Sprintf("-%v", int(data.Main.TemperatureFeels))
	default:
		feelsTemp = fmt.Sprintf("%v", int(data.Main.TemperatureFeels))
	}

	reply += fmt.Sprintf("\nТемпература повітря %s С", temp)
	reply += fmt.Sprintf("\nМін: %s С Макс: %s C", minTemp, maxTemp)
	reply += fmt.Sprintf("\nВідчувається наче %s С", feelsTemp)

	reply += fmt.Sprintf("\nВологість повітря %d %%", data.Main.Humidity)

	reply += fmt.Sprintf("\n%s", data.Weather[0].Description)

	reply += fmt.Sprintf("\nШвидкість вітру %v км/год", data.Wind.Speed)

	return
}

func createURL(city string) (*url.URL, error) {
	appID := "db9a441fce153ac5701b2235510e4d1b"

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
	resp, err := http.Get("https://" + url.String())
	if err != nil {
		return nil, fmt.Errorf("sendRequest Get error %w", err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
