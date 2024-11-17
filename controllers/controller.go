package controllers

import (
	auth_controller "rms-api/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth", auth_controller.Auth)

	api.Get("/hc", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
}
