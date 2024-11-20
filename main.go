package main

import (
	"log"
	"rms-api/controllers"
	"rms-api/db"
	"rms-api/redis"
	"rms-api/services/validator"
	"rms-api/telegram"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	controllers.Setup(app)
	validator.Init()

	redis.Connect()
	db.Connect()
	go telegram.Start()

	log.Fatal(app.Listen(":3000"))
}
