package repositories

import (
	"encoding/json"
	"errors"
	"github.com/levigross/grequests"
	"hexago/internal/application/dtos"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://api.cryptapi.io/"
)

type CoinRepository interface {
	GetBySymbol(symbol string) (dtos.CoinResponseDTO, error)
}

type CoinGeckoClient struct {
}

func NewCoinRepository() *CoinGeckoClient {
	return &CoinGeckoClient{}
}

func (c *CoinGeckoClient) GetBySymbol(symbol string) (dtos.CoinResponseDTO, error) {
	var coinResponse dtos.CoinResponseDTO

	response, err := grequests.Get(baseUrl+strings.ToLower(symbol)+"/info", nil)
	if err != nil {
		return coinResponse, err
	}

	if response.StatusCode >= http.StatusBadRequest {
		err = json.Unmarshal(response.Bytes(), &coinResponse)
		if err != nil {
			return dtos.CoinResponseDTO{}, err
		}
		return dtos.CoinResponseDTO{}, errors.New(coinResponse.Error)
	}

	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusNotFound {
		err = json.Unmarshal(response.Bytes(), &coinResponse)
		if err != nil {
			return coinResponse, err
		}
	}

	return coinResponse, nil
}
