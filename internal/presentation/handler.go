package presentation

import (
	"encoding/json"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/application"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/presentation/validation"
	"io"
	"net/http"
)

type zipcodeRequest struct {
	Zipcode string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
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

	zipcodeService, err := application.NewZipcodeService(application.Zipcode(zipcodeRequest.Zipcode))
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

	weatherService := application.NewWeatherService(place)
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
