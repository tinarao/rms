package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	app := fiber.New()

	app.Get("/hc", func(c fiber.Ctx) error {
		return c.SendString("Ok")
	})

	log.Fatal(app.Listen(":3000"))
}
