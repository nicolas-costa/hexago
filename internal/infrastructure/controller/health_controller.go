package controller

import "github.com/gofiber/fiber/v2"

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Check(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"healthy": true,
	})
}
