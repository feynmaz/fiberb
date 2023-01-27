package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"

	"github.com/feynmaz/fiberg/configs"
	"github.com/feynmaz/fiberg/handlers"
	"github.com/feynmaz/fiberg/utils"
)

func main() {
	config := configs.GetDefault()
	utils.InitLogger(config.Debug)

	app := fiber.New()

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "{\"level\":\"info\",\"time\":\"${time}\",\"request_id\":\"${locals:requestid}\",\"status\":${status},\"method\":\"${method}\",\"path\":\"${path}\"}\n",
		TimeFormat: time.RFC3339,
	}))

	handler := handlers.NewHandler()

	app.Get("/healthcheck", handler.HealthCheck)

	app.Use(handlers.NotFound)

	zap.L().With(
		zap.Error(
			app.Listen(fmt.Sprintf(":%d", config.Port)),
		),
	).Fatal("")
}
