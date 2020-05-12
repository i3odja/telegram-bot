package weather

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"../../../model"
	"../helper"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	weatherHost           = "api.openweathermap.org/data/2.5/weather"
	weatherCaption        = "Ваш прогноз погоди готовий!\n\n"
	weatherWarningMessage = "‼️ упс... ‼️ Щось пішло не так. Сервіс прогнозу погоди працює не корекно. Я виправлю це найближчим часом ❌"

	messageWeatherCity              = "Місто \t%s"
	messageWeatherOnToday           = "\nПрогноз погоди на сьогодні %v-%v-%v"
	messageWeatherTemperature       = "\n🌡 Температура повітря %s °C"
	messageWeatherTemperatureMinMax = "\n🌡 Мін: %s °C Макс: %s °C"
	messageWeatherTemperatureFeels  = "\n🌡 Відчувається наче %s °C"
	messageWeatherHumidity          = "\nВологість повітря %d %%"
	messageWeatherDescription       = "\n%s"
	messageWeatherWindSpeed         = "\n🌬 Швидкість вітру %v км/год"

	weatherTemperatureBelowZero = "+%v"

	weatherUnits = "metric"
)

func Forecast(city string, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	// TODO: Delete before send to GitHub!!!
	os.Setenv("APP_ID", "db9a441fce153ac5701b2235510e4d1b")
	appID := os.Getenv("APP_ID")
	if appID == "" {
		warningMsg := fmt.Sprintf(weatherWarningMessage)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, warningMsg)

		_, err := bot.Send(msg)
		if err != nil {
			return fmt.Errorf("forecast Getenv Send error %w", err)
		}

		return fmt.Errorf("sorry, but you did not setup APP_ID, please fix it")
	}

	dataWeather, err := jsonParse(appID, city, weatherUnits)
	if err != nil {
		return err
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, weatherCaption))
	if err != nil {
		return fmt.Errorf("forecast Send error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, makeReplyWeather(dataWeather)))
	if err != nil {
		return fmt.Errorf("forecast CreateMessage Send error %w", err)
	}

	return nil
}

// jsonParse gets URL by using createURL(),
// gets JSON body from request by using sendRequest(),
// decodes received JSON and returns *model.DataWeather type output data
// or error if something was wrong
func jsonParse(appID string, city string, units string) (*model.DataWeather, error) {
	u, err := createWeatherURL(appID, city, units)
	if err != nil {
		return nil, fmt.Errorf("jsonParse createURL error %w", err)
	}

	ju, err := helper.SendRequest(u)
	if err != nil {
		return nil, fmt.Errorf("jsonParse sendRequest error %w", err)
	}

	data := new(model.DataWeather)
	err = json.Unmarshal(ju, &data)
	if err != nil {
		return nil, fmt.Errorf("jsonParse JSON Unmarshal error %w", err)
	}

	return data, nil
}

// createURL passes APPID, city and units which will be used as URL parameters
// and returns URL or error if something was wrong
//
//	Example:
//		exapmleURL := "example.com.ua"
//
//		url, err := createURL("myAppID", "Lviv")
//		if err != nil {
//			return err
//		}
//
//		fmt.Println(url)
//
//	Output:
//		example.com.ua?APPID=myAppID&city=Lviv&units=metric
func createWeatherURL(appID string, city string, units string) (*url.URL, error) {
	u, err := url.Parse(weatherHost)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url error: %w", err)
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url query error: %w", err)
	}

	q.Set("APPID", appID)
	q.Set("q", city)
	q.Set("units", units)

	u.RawQuery = q.Encode()

	return u, nil
}

func makeReplyWeather(data *model.DataWeather) (reply string) {
	y, m, d := time.Now().Date()

	reply += fmt.Sprintf(messageWeatherCity, data.Name)
	reply += fmt.Sprintf(messageWeatherOnToday, d, m, y)
	reply += fmt.Sprintf(messageWeatherTemperature, signOfTemperature(data.Main.Temperature))
	reply += fmt.Sprintf(messageWeatherTemperatureMinMax, signOfTemperature(data.Main.TemperatureMin), signOfTemperature(data.Main.TemperatureMax))
	reply += fmt.Sprintf(messageWeatherTemperatureFeels, signOfTemperature(data.Main.TemperatureFeels))
	reply += fmt.Sprintf(messageWeatherHumidity, data.Main.Humidity)
	reply += fmt.Sprintf(messageWeatherDescription, data.Weather[0].Description)
	reply += fmt.Sprintf(messageWeatherWindSpeed, data.Wind.Speed)

	return
}

func signOfTemperature(degree float32) string {
	if int(degree) > 0 {
		return fmt.Sprintf(weatherTemperatureBelowZero, int16(degree))
	}

	return fmt.Sprintf("%v", int(degree))
}
