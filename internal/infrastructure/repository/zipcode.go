package repository

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/domain/repository"
)

func NewZipcodeRepository() repository.ZipcodeRepository {
	return NewViaCep()
}
