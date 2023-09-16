package controller

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

func (c *CoinController) CheckPrice(ctx *fiber.Ctx) error {
	coinData, err := c.service.GetCurrentPriceBySymbol(ctx.Query("symbol", "ETH"))
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(coinData)
}
