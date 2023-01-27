package handlers

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendString("Not Found")
}
