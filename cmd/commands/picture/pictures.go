package picture

import (
	"fmt"
	"math/rand"
	"net/url"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	pictureURL = "https://picsum.photos/600/600?random"

	randomMax = 1000
	randomMin = 1
)

func Picture(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	picture, err := createPictureURL()
	if err != nil {
		return fmt.Errorf("Picture createURL error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ваша картинка готова!"))
	if err != nil {
		return fmt.Errorf("Picture Send message error %w", err)
	}

	_, err = bot.Send(tgbotapi.NewPhotoShare(update.Message.Chat.ID, picture.String()))
	if err != nil {
		return fmt.Errorf("Picture NewPhotoShare error %w", err)
	}

	return nil
}

func createPictureURL() (*url.URL, error) {
	u, err := url.Parse(pictureURL)
	if err != nil {
		return nil, fmt.Errorf("createPictureURL parse url error: %w", err)
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("createPictureURL parse url query error: %w", err)
	}

	q.Set("random", string(rand.Intn((randomMax-randomMin)+randomMin)))

	u.RawQuery = q.Encode()

	return u, nil
}
