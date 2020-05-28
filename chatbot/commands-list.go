package chatbot

import (
	"fmt"

	//"../cmd/commands/covid"
	//"../cmd/commands/currency"
	//"../cmd/commands/greeter"
	//"../cmd/commands/joke"
	//"../cmd/commands/picture"
	//"../cmd/commands/weather"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	"github.com/i3odja/telegram-bot/cmd/commands/covid"
	"github.com/i3odja/telegram-bot/cmd/commands/currency"
	"github.com/i3odja/telegram-bot/cmd/commands/greeter"
	"github.com/i3odja/telegram-bot/cmd/commands/joke"
	"github.com/i3odja/telegram-bot/cmd/commands/picture"
	"github.com/i3odja/telegram-bot/cmd/commands/weather"
)

func SelectCommandsList(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	switch update.Message.Command() {
	case "hello":
		err := greeter.Greeter(bot, *update)
		return checkError(err)
	case "weather":
		city := update.Message.CommandArguments()
		if city == "" {
			city = "Lviv"
		}
		err := weather.Forecast(city, bot, update)
		return checkError(err)
	case "joke":
		err := joke.Joke(bot, update)
		return checkError(err)
	case "picture":
		err := picture.Picture(bot, update)
		return checkError(err)
	case "currency":
		err := currency.Currency(bot, update)
		return checkError(err)
	case "covid":
		err := covid.Statistic(bot, update)
		return checkError(err)
	case "help":
		err := CommandList(bot, update)
		return checkError(err)
	case "buttons":
		err := KeyboardButtons(bot, update)
		return checkError(err)
	default:
		err := replyUnknownCommand(bot, update, "I'm waiting for your commands. I will help you!")
		return checkError(err)
	}
}

func CommandList(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	help := fmt.Sprintf("Hello, I'm your personal bot!\n")
	help += fmt.Sprintf("\n")

	help += fmt.Sprintf("Please, use the following commands:\n")
	help += fmt.Sprintf("\n")

	help += fmt.Sprintf("/buttons - ☑️ I will show you all commands like buttons!\n")
	help += fmt.Sprintf("/help - ❓ I will show you all available commands!\n")
	help += fmt.Sprintf("\n")

	help += fmt.Sprintf("/hello - 👋 I want to say hello to you!\n")
	help += fmt.Sprintf("/weather - 🌤 I want to show you the weather! By default I will show forecast for Lviv. If You want to get forecast for another city, please use the next format:\n")
	help += fmt.Sprintf("    /weather <city>\n")
	help += fmt.Sprintf("    for example: /weather Kyiv or /weather Київ\n")

	help += fmt.Sprintf("/joke - 😂 I want to show you very funny Joke!\n")
	help += fmt.Sprintf("/picture - 🌄 I want to show you very interesting Picture!\n")
	help += fmt.Sprintf("/currency - 💰 I want to show you the current currency!\n")
	help += fmt.Sprintf("/covid - \U0001F9A0 I want to show all cases of covid-19 on yesterday!\n")

	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, help))
	if err != nil {
		return fmt.Errorf("replyUnknownCommand Send error %w", err)
	}

	return nil
}

func replyUnknownCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string) error {
	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
	if err != nil {
		return fmt.Errorf("replyUnknownCommand Send error %w", err)
	}

	return nil
}

func checkError(err error) error {
	if err != nil {
		return fmt.Errorf("SelectCommandsList error %w", err)
	}

	return nil
}
