package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
)

type WeatherRepositoryMock struct {
}

func NewWeatherRepositoryMock() repository.WeatherRepository {
	return &WeatherRepositoryMock{}
}

func (w *WeatherRepositoryMock) FetchWeather(p entity.Place) (*entity.WeatherResponse, error) {
	return entity.NewWeatherResponse("São Paulo", 25, 77, 298), nil
}
