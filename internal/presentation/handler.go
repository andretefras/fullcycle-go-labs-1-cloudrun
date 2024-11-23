package presentation

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, validation.ErrValidatingRequestMethod, validation.ErrorCodes[validation.ErrValidatingRequestMethod])
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s\n", string(body))
		http.Error(w, validation.ErrReadingRequestBody, validation.ErrorCodes[validation.ErrReadingRequestBody])
		return
	}
	defer r.Body.Close()

	var zipcodeRequest zipcodeRequest
	err = json.Unmarshal(body, &zipcodeRequest)
	if err != nil {
		http.Error(w, validation.ErrValidatingZipcode, validation.ErrorCodes[validation.ErrValidatingZipcode])
		return
	}

	zipcodeService, err := application.NewZipcodeService(application.Zipcode(zipcodeRequest.Zipcode))
	if err != nil {
		http.Error(w, err.Error(), validation.ErrorCodes[err.Error()])
		return
	}

	if zipcodeService == nil {
		http.Error(w, validation.ErrValidatingZipcode, validation.ErrorCodes[validation.ErrValidatingZipcode])
		return
	}

	place, err := zipcodeService.GetPlace()
	if err != nil {
		http.Error(w, err.Error(), validation.ErrorCodes[err.Error()])
		return
	}

	weatherService := application.NewWeatherService(place)
	weather, err := weatherService.GetWeather()
	if err != nil {
		http.Error(w, err.Error(), validation.ErrorCodes[err.Error()])
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		http.Error(w, validation.ErrRequestingWeather, validation.ErrorCodes[validation.ErrRequestingWeather])
		return
	}
}
