package application

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/infrastructure/repository"
)

type Weather struct {
	City       string  `json:"city"`
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
	Kelvin     float64 `json:"temp_k"`
}

func NewWeather(city string, celsius float64, fahrenheit float64, kelvin float64) *Weather {
	return &Weather{City: city, Celsius: celsius, Fahrenheit: fahrenheit, Kelvin: kelvin}
}

func NewWeatherService(p *Place) WeatherService {
	return WeatherService{place: p}
}

type WeatherService struct {
	place *Place
}

func (s *WeatherService) GetWeather() (*Weather, error) {
	weatherRepository := repository.NewWeatherRepository()
	weather, err := weatherRepository.FetchWeather(entity.Place(*s.place))
	if err != nil {
		return nil, err
	}
	return NewWeather(
		weather.City,
		weather.Celsius,
		weather.Fahrenheit,
		weather.Kelvin,
	), nil
}
