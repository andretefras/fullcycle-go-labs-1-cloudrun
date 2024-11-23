package validation

import "net/http"

const (
	ErrValidatingRequestMethod = "can not allow method"
	ErrReadingRequestBody      = "can not read request"
	ErrValidatingZipcode       = "invalid zipcode"
	ErrRequestingZipcode       = "can not request zipcode"
	ErrFindingZipcode          = "can not find zipcode"
	ErrReadingZipcode          = "can not read zipcode"
	ErrParsingZipcode          = "can not parse zipcode"
	ErrRequestingWeather       = "can not request weather"
	ErrMissingApiKey           = "can not find weather api key"
	ErrParsingWeather          = "can not parse weather"
)

var ErrorCodes = map[string]int{
	ErrValidatingRequestMethod: http.StatusMethodNotAllowed,
	ErrReadingRequestBody:      http.StatusBadRequest,
	ErrValidatingZipcode:       http.StatusUnprocessableEntity,
	ErrRequestingZipcode:       http.StatusUnprocessableEntity,
	ErrFindingZipcode:          http.StatusNotFound,
	ErrReadingZipcode:          http.StatusNotFound,
	ErrParsingZipcode:          http.StatusAccepted,
	ErrRequestingWeather:       http.StatusNotFound,
	ErrMissingApiKey:           http.StatusBadRequest,
	ErrParsingWeather:          http.StatusNotFound,
}
