package main

import (
	"example/internal/config"
	"log"
	"os"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't create zap logger %s", err)
		os.Exit(1)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Fatal("Can't sync logger after shutdown", zap.Error(err))
		}
	}()

	sugar := logger.Sugar()

	cfg := config.ReadConfig()

	goose.SetLogger(zap.NewStdLog(logger))

	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
		},
	)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	sugar.Fatal(app.Listen(cfg.Addr))
}
