package chatbot_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/i3odja/telegram-bot/chatbot"
	"github.com/i3odja/telegram-bot/cmd/commands/greeter"
	"github.com/i3odja/telegram-bot/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func TestSetupUserInfo(t *testing.T) {
	user := &tgbotapi.User{
		ID:           11111,
		FirstName:    "FirstName",
		LastName:     "LastName",
		UserName:     "chatBot_1",
		LanguageCode: "uk",
		IsBot:        true,
	}

	expected := &model.User{
		ID:           11111,
		Login:        "chatBot_1",
		FirstName:    "FirstName",
		LastName:     "LastName",
		LanguageCode: "uk",
	}

	get, err := greeter.SetupUserInfo(user)
	require.NoError(t, err)

	assert.Equal(t, expected, get)
}

func TestCreateNewBotConnectionError(t *testing.T) {
	_, err := chatbot.CreateNewBotConnection()
	require.Error(t, err)

	assert.EqualError(t, err, "wrong telegram token =(")
}

func TestCreateNewBotConnectionUnauthorized(t *testing.T) {
	os.Setenv("TOKEN_TG_BOT", "1101236908:wrong-token")
	bot, err := chatbot.CreateNewBotConnection()
	require.Error(t, err)

	assert.Nil(t, bot)

	assert.EqualError(t, err, "Unauthorized")
}

func TestCreateNewBotConnectionSuccess(t *testing.T) {
	os.Setenv("TOKEN_TG_BOT", "1101236908:AAEgW4h902L8rceydRd0FmF3kMPWnCTNEaw")
	bot, err := chatbot.CreateNewBotConnection()
	require.NoError(t, err)

	assert.NotNil(t, bot)
}

func TestCreateReply(t *testing.T) {
	tests := []struct {
		name     string
		user     *model.User
		expected string
		message  string
	}{
		{
			name: "test1",
			user: &model.User{
				ID:           777,
				Login:        "user1",
				FirstName:    "Андрій",
				LastName:     "Шевченко",
				LanguageCode: "uk",
			},
			expected: "Привіт Андрій! Як у тебе справи?",
			message:  "Ukraine!",
		},
		{
			name: "test2",
			user: &model.User{
				ID:           777,
				Login:        "user1",
				FirstName:    "Cristiano",
				LastName:     "Ronaldo",
				LanguageCode: "en",
			},
			expected: "Hello Cristiano! How are you?",
			message:  "English!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := greeter.CreateReply(tt.user)
			if get != tt.expected {
				t.Errorf("want %v but got %v", tt.expected, get)
			}
		})
	}
}

func TestSetupToken(t *testing.T) {
	os.Setenv("TOKEN_TG_BOT", "1101236908:AAGdRKCvt8EzpByAFjPKnof-gYKjdTE9jVM")

	token, err := chatbot.SetupToken()
	require.NoError(t, err)
	assert.Equal(t, "1101236908:AAGdRKCvt8EzpByAFjPKnof-gYKjdTE9jVM", token)
}
