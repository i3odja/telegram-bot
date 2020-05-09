package model

type Covid19 struct {
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Province    string `json:"province"`
	City        string `json:"city"`
	CityCode    string `json:"city_code"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Cases       int    `json:"cases"`
	Status      string `json:"status"`
	Date        string `json:"date"`
}
