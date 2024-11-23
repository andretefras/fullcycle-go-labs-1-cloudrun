package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
	zipcodeerrors "github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/presentation/validation"
	"io"
	"net/http"
	"net/url"
	"os"
)

type WeatherApi struct {
}

func NewWeatherApi() repository.WeatherRepository {
	return &WeatherApi{}
}

func (w *WeatherApi) FetchWeather(p entity.Place) (*entity.WeatherResponse, error) {
	params := url.Values{}
	params.Add("q", string(p))
	fullUrl := fmt.Sprintf("%s?%s", "https://api.weatherapi.com/v1/current.json", params.Encode())
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, errors.New(zipcodeerrors.ErrRequestingWeather)
	}

	weatherApiKey, ok := os.LookupEnv("WEATHER_API_KEY")
	if !ok {
		return nil, errors.New(zipcodeerrors.ErrMissingApiKey)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("key", weatherApiKey)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(zipcodeerrors.ErrRequestingWeather)
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(zipcodeerrors.ErrRequestingWeather)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(zipcodeerrors.ErrRequestingWeather)
	}
	fmt.Printf("%s\n", string(body))

	var weatherApiResponse map[string]interface{}
	err = json.Unmarshal(body, &weatherApiResponse)
	if err != nil {
		return nil, errors.New(zipcodeerrors.ErrParsingWeather)
	}

	return entity.NewWeatherResponse(
		weatherApiResponse["location"].(map[string]interface{})["name"].(string),
		weatherApiResponse["current"].(map[string]interface{})["temp_c"].(float64),
		weatherApiResponse["current"].(map[string]interface{})["temp_f"].(float64),
		weatherApiResponse["current"].(map[string]interface{})["temp_c"].(float64)+273,
	), nil
}
