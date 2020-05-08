package commands

import (
	"fmt"
	"math/rand"
	"net/url"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const pictureURL = "https://picsum.photos/600/600?random"

func getPicture(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	picture, err := createPictureURL()
	if err != nil {
		return fmt.Errorf("getPicture createURL error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ваша картинка готова!"))
	if err != nil {
		return fmt.Errorf("getPicture Send message error %w", err)
	}

	msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, picture.String())
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("getPicture NewPhotoShare error %w", err)
	}

	return nil
}

func createPictureURL() (*url.URL, error) {
	min, max := 1, 1000
	number := string(rand.Intn((max - min) + min))

	u, err := url.Parse(pictureURL)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url error: %w", err)
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("createURL parse url query error: %w", err)
	}

	q.Set("random", number)

	u.RawQuery = q.Encode()

	return u, nil
}
