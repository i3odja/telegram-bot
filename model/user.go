package model

type User struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}
