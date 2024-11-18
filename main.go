package main

import (
	"log"
	"rms-api/controllers"
	"rms-api/db"
	"rms-api/redis"
	"rms-api/services/validator"
	"rms-api/telegram"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())
	controllers.Setup(app)
	validator.Init()

	redis.Connect()
	db.Connect()
	go telegram.Start()

	log.Fatal(app.Listen(":3000"))
}
