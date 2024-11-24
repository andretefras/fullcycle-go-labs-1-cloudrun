package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
)

const (
	ZipcodeMockKey   = "mock"
	ZipcodeViacepKey = "viacep"
)

func NewZipcodeRepository(r string) repository.ZipcodeRepository {
	if r == ZipcodeMockKey {
		return NewZipcodeRepositoryMock()
	}
	return NewViaCep()
}
