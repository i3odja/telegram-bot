package joke

import (
	"encoding/json"
	"fmt"
	"net/url"

	//"../../../model"
	//"../helper"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	"github.com/i3odja/telegram-bot/cmd/commands/helper"
	"github.com/i3odja/telegram-bot/model"
)

const jokesURL = "api.chucknorris.io/jokes/random"

func Joke(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	urlJoke, err := url.Parse(jokesURL)
	if err != nil {
		return fmt.Errorf("joke Parse error %w", err)
	}

	jsonJoke, err := helper.SendRequest(urlJoke)
	if err != nil {
		return fmt.Errorf("joke sendRequest error %w", err)
	}

	// Parse JSON
	dataJoke := new(model.Joke)
	err = json.Unmarshal(jsonJoke, &dataJoke)
	if err != nil {
		return fmt.Errorf("joke JSON Unmarshal error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш анекдот готовий!"))
	if err != nil {
		return fmt.Errorf("joke Send error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, dataJoke.Value))
	if err != nil {
		return fmt.Errorf("joke Send error %w", err)
	}

	return nil
}
