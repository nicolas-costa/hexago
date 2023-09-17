package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hexago/internal/application"
)

type CoinController struct {
	service application.Pricer
}

func NewCoinController(coinService application.Pricer) *CoinController {
	return &CoinController{service: coinService}
}

// Get crypto price by symbol
//
//	@Summary      Get crypto
//	@Description  get accounts
//	@Tags         crypto
//	@Accept       json
//	@Produce      json
//	@Param        symbol    query     string  true  "get data by coin symbol"
//	@Success      200  {object}   dtos.CoinResponseDTO
//	@Failure      404
//	@Failure      500
//	@Router       /coin [get]
func (c *CoinController) CheckPrice(ctx *fiber.Ctx) error {
	coinData, err := c.service.GetCurrentPriceBySymbol(ctx.Query("symbol", "ETH"))
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(coinData)
}
