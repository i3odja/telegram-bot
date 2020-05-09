package model

type DataWeather struct {
	Coordinates struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}
	Base string `json:"base"`
	Main struct {
		Temperature      float32 `json:"temp"`
		TemperatureFeels float32 `json:"feels_like"`
		TemperatureMin   float32 `json:"temp_min"`
		TemperatureMax   float32 `json:"temp_max"`
		Pressure         int     `json:"pressure"`
		Humidity         int     `json:"humidity"`
	}
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	}
	Clouds struct {
		All int `json:"all"`
	}
	DT  int `json:"dt"`
	SYS struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	}
	TimeZone int    `json:"time_zone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
