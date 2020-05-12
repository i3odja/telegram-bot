package currency_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"../../model"
	"../commands"
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
			expected: "USD\nСтаном на 12.12.2012\nКурс: 123\n",
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
			expected: "EUR\nСтаном на 12.12.2012\nКурс: 123\n",
			msg:      "Correct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply := commands.MakeReplyCurrency(tt.actual)

			assert.Equal(t, tt.expected, reply[0])
		})
	}
}
