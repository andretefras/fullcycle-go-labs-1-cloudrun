package entity

type WeatherResponse struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
	Kelvin     float64 `json:"temp_k"`
}

func NewWeatherResponse(city string, celsius float64, fahrenheit float64, kelvin float64) *WeatherResponse {
	return &WeatherResponse{City: city, Celsius: celsius, Fahrenheit: fahrenheit, Kelvin: kelvin}
}
