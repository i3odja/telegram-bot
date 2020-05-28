package currency_test

import (
	"testing"

	"github.com/i3odja/telegram-bot/cmd/commands/currency"
	"github.com/i3odja/telegram-bot/model"

	"github.com/stretchr/testify/assert"
	//"../../../model"
	//"../../commands/currency"
)

func TestMakeReplyCurrency(t *testing.T) {
	tests := []struct {
		name     string
		actual   []model.Currency
		expected string
		msg      string
	}{
		{
			name: "Correct result",
			actual: []model.Currency{
				{
					R030:         111,
					TXT:          "USD",
					Rate:         123,
					CC:           "USD",
					Exchangedate: "12.12.2012",
				},
			},
			expected: "🇺🇸 USD\nКурс: 123\n\n",
			msg:      "Correct",
		}, {
			name: "Correct result",
			actual: []model.Currency{
				{
					R030:         111,
					TXT:          "EUR",
					Rate:         123,
					CC:           "EUR",
					Exchangedate: "12.12.2012",
				},
			},
			expected: "🇪🇺 EUR\nКурс: 123\n\n",
			msg:      "Correct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := currency.MakeReplyCurrency(tt.actual)

			assert.Equal(t, tt.expected, reply[0])
		})
	}
}

func TestMakeReplyCurrencyInvalid(t *testing.T) {
	testCase := struct {
		name     string
		actual   []model.Currency
		expected string
		msg      string
	}{

		name: "Invalid result",
		actual: []model.Currency{
			{
				R030:         111,
				TXT:          "USD",
				Rate:         123,
				CC:           "USD",
				Exchangedate: "12.12.2012",
			},
		},
		expected: "🇺🇸 USD\nКурс: 999\n\n",
		msg:      "Incorrect",
	}

	t.Run(testCase.name, func(t *testing.T) {
		reply := currency.MakeReplyCurrency(testCase.actual)

		assert.NotEqual(t, testCase.expected, reply)
	})
}
