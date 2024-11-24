package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/entity"
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
)

func NewZipcodeRepositoryMock() repository.ZipcodeRepository {
	return &Mock{}
}

type Mock struct {
}

func (*Mock) FetchZipcode(z entity.Zipcode) (*entity.ZipcodeResponse, error) {
	return entity.NewZipcodeResponse("São Paulo", ""), nil
}
