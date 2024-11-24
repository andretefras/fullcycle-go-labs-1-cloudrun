package application

import (
	"errors"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	rinterface "github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/infrastructure/repository"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/presentation/validation"
)

type Zipcode string

type Place string

func NewZipcodeService(zipcode Zipcode, r string) (*ZipcodeService, error) {
	if len(zipcode) != 8 {
		return nil, errors.New(validation.ErrValidatingZipcode)
	}
	return &ZipcodeService{
		zipcode:    zipcode,
		repository: repository.NewZipcodeRepository(r),
	}, nil
}

type ZipcodeService struct {
	zipcode    Zipcode
	repository rinterface.ZipcodeRepository
}

func (s *ZipcodeService) GetPlace() (*Place, error) {
	zipcode, err := s.repository.FetchZipcode(entity.Zipcode(s.zipcode))
	if err != nil {
		return nil, err
	}
	place := Place(zipcode.Localidade)
	return &place, nil
}
