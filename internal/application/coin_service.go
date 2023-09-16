package application

import (
	"hexago/internal/application/dtos"
	"hexago/internal/infrastructure/repositories"
)

type Pricer interface {
	GetCurrentPriceBySymbol(symbol string) (dtos.CoinResponseDTO, error)
}

type CoinService struct {
	repository repositories.CoinRepository
}

func NewCoinService(coinRepository repositories.CoinRepository) *CoinService {
	return &CoinService{repository: coinRepository}
}

func (c *CoinService) GetCurrentPriceBySymbol(symbol string) (dtos.CoinResponseDTO, error) {
	return c.repository.GetBySymbol(symbol)
}
