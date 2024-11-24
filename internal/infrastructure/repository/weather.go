package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
)

const (
	WeatherMockKey = "mock"
	WeatherApiKey  = "weatherapi"
)

func NewWeatherRepository(r string) repository.WeatherRepository {
	if r == WeatherMockKey {
		return NewWeatherRepositoryMock()
	}
	return NewWeatherApi()
}
