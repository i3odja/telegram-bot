package commands

import (
	"fmt"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func SelectCommandsList(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	switch update.Message.Command() {
	case "hello":
		err := greeter(bot, *update)
		return checkError(err)
	case "weather":
		city := update.Message.CommandArguments()
		if city == "" {
			city = "Lviv"
		}
		err := getWeather(city, bot, update)
		return checkError(err)
	case "joke":
		err := getJoke(bot, update)
		return checkError(err)
	case "picture":
		err := getPicture(bot, update)
		return checkError(err)
	case "currency":
		err := Currency(bot, update)
		return checkError(err)
	case "covid":
		err := Covid19(bot, update)
		return checkError(err)
	case "help":
		err := GetCommandList(bot, update)
		return checkError(err)
	default:
		err := replyForCommands(bot, update, "I'm waiting for your commands. I will help you!")
		return checkError(err)
	}
}

func GetCommandList(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	help := fmt.Sprintf("/help - I will show you all available commands!\n")
	help += fmt.Sprintf("/hello - I want to say hello to you!\n")
	help += fmt.Sprintf("/weather - I want to show you the weather!\n")
	help += fmt.Sprintf("/joke - I want to show you very funny joke!\n")
	help += fmt.Sprintf("/picture - I want to show you very interesting picture!\n")
	help += fmt.Sprintf("/currency - I want to show you the current currency!\n")
	help += fmt.Sprintf("/covid - I want to show all cases of covid-19 on yesterday!\n")

	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, help))
	if err != nil {
		return fmt.Errorf("replyForCommands Send error %w", err)
	}

	return nil
}

func replyForCommands(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string) error {
	_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
	if err != nil {
		return fmt.Errorf("replyForCommands Send error %w", err)
	}

	return nil
}

func checkError(err error) error {
	if err != nil {
		return fmt.Errorf("SelectCommandsList %w", err)
	}

	return nil
}
