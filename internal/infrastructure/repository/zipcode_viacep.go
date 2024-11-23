package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
	validationerrors "github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/presentation/validation"
	"io"
	"net/http"
)

func NewViaCep() repository.ZipcodeRepository {
	return &ViaCep{}
}

type ViaCep struct {
}

func (*ViaCep) FetchZipcode(z entity.Zipcode) (*entity.ZipcodeResponse, error) {
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+string(z)+"/json/", nil)
	if err != nil {
		return nil, errors.New(validationerrors.ErrRequestingZipcode)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(validationerrors.ErrFindingZipcode)
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(validationerrors.ErrFindingZipcode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(validationerrors.ErrRequestingZipcode)
	}
	fmt.Printf("%s\n", string(body))

	var zipcodeResponse entity.ZipcodeResponse
	err = json.Unmarshal(body, &zipcodeResponse)
	if err != nil {
		return nil, errors.New(validationerrors.ErrParsingZipcode)
	}

	if zipcodeResponse.Erro != "" {
		return nil, errors.New(validationerrors.ErrFindingZipcode)
	}

	return &zipcodeResponse, nil
}
