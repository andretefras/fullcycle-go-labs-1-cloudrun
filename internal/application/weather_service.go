package application

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	rpi "github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
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

func NewWeatherService(p *Place, r string) WeatherService {
	return WeatherService{
		place:      p,
		repository: repository.NewWeatherRepository(r),
	}
}

type WeatherService struct {
	place      *Place
	repository rpi.WeatherRepository
}

func (s *WeatherService) GetWeather() (*Weather, error) {
	weather, err := s.repository.FetchWeather(entity.Place(*s.place))
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
