package controllers

import (
	adminauth_controller "rms-api/controllers/admin-auth"
	auth_controller "rms-api/controllers/auth"
	restaurants_controller "rms-api/controllers/restaurants"
	"rms-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth", auth_controller.Auth)

	api.Get("/hc", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	api.Post("/aregister", adminauth_controller.Register)
	api.Post("/alogin", adminauth_controller.Login)
	api.Get("/averify", adminauth_controller.Verify)

	api.Get("/restaurants/all", restaurants_controller.GetAllRestaurants)
	api.Get("/restaurants/slug/:slug", restaurants_controller.GetRestaurantBySlug)

	// Protected
	protected := api.Group("/p", middleware.AdminOnly)
	protected.Post("/restaurants/create", restaurants_controller.CreateRestaurant)
}
