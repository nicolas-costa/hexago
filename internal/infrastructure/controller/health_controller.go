package controller

import (
	"github.com/gofiber/fiber/v2"
	"hexago/internal/application"
)

type HealthController struct {
	healthService application.Checker
}

func NewHealthController(service application.Checker) *HealthController {
	return &HealthController{
		healthService: service,
	}
}

func (h *HealthController) Check(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"healthy": h.healthService.Check(),
	})
}
