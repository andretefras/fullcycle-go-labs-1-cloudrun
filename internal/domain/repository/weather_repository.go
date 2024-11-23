package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
)

type WeatherRepository interface {
	FetchWeather(p entity.Place) (*entity.WeatherResponse, error)
}
