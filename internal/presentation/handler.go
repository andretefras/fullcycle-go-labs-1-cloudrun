package presentation

import (
	"encoding/json"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/application"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/infrastructure/repository"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/presentation/validation"
	"io"
	"net/http"
	"os"
)

type zipcodeRequest struct {
	Zipcode string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		httpError(w, validation.ErrValidatingRequestMethod)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpError(w, validation.ErrReadingRequestBody)
		return
	}
	defer r.Body.Close()

	var zipcodeRequest zipcodeRequest
	err = json.Unmarshal(body, &zipcodeRequest)
	if err != nil {
		httpError(w, validation.ErrValidatingZipcode)
		return
	}

	zipcodeRepository := os.Getenv("ZIPCODE_REPOSITORY")
	if zipcodeRepository == "" {
		zipcodeRepository = repository.ZipcodeMockKey
	}

	zipcodeService, err := application.NewZipcodeService(application.Zipcode(zipcodeRequest.Zipcode), zipcodeRepository)
	if err != nil {
		http.Error(w, err.Error(), validation.ErrorCodes[err.Error()])
		return
	}

	if zipcodeService == nil {
		httpError(w, validation.ErrValidatingZipcode)
		return
	}

	place, err := zipcodeService.GetPlace()
	if err != nil {
		httpError(w, err.Error())
		return
	}

	weatherRepository := os.Getenv("WEATHER_REPOSITORY")
	if weatherRepository == "" {
		weatherRepository = repository.WeatherMockKey
	}

	weatherService := application.NewWeatherService(place, weatherRepository)
	weather, err := weatherService.GetWeather()
	if err != nil {
		httpError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		httpError(w, validation.ErrRequestingWeather)
		return
	}
}

func httpError(w http.ResponseWriter, e string) {
	http.Error(w, e, validation.ErrorCodes[e])
}
