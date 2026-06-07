package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Get("/version", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    "healthcheck-rest-api",
			"version": "1.0.0",
			"status":  "ok",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
