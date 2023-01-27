package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func (h Handler) HealthCheck(c *fiber.Ctx) error {
	requestID := c.Locals(requestid.ConfigDefault.ContextKey).(string)
	zap.L().With(
		zap.String("request_id", requestID),
	).Debug("inside HealthCheck")
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
