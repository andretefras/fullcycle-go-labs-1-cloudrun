package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
)

type ZipcodeRepository interface {
	FetchZipcode(z entity.Zipcode) (*entity.ZipcodeResponse, error)
}
