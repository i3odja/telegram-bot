package chatbot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"../model"
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

	get, err := SetupUserInfo(user)
	require.NoError(t, err)

	assert.Equal(t, expected, get)
}

func TestCreateNewBotConnectionSuccess(t *testing.T) {
	bot, err := CreateNewBotConnection("1161561075:AAG6WNCUAgAH0V-l5CG2QGo5smCzELERSow")
	require.NoError(t, err)

	assert.NotNil(t, bot)
}

func TestCreateNewBotConnectionError(t *testing.T) {
	_, err := CreateNewBotConnection("21161561075:AAG6WNCUAgAH0V-l5CG2QGo5smCzELERSow")

	assert.EqualError(t, fmt.Errorf("CreateNewBotConnection error Unauthorized"), err.Error())
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
			get := CreateReply(tt.user)
			if get != tt.expected {
				t.Errorf("want %v but got %v", tt.expected, get)
			}
		})
	}
}
