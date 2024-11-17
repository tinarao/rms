package main

import (
	"log"
	"os"
	"rms-api/controllers"
	"rms-api/db"
	"rms-api/redis"
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

	redisPassword := os.Getenv("REDIS_PASSWORD")
	app := fiber.New()
	app.Use(cors.New())
	controllers.Setup(app)
	redis.Init(redisPassword)
	db.Init()
	go telegram.Init()

	log.Fatal(app.Listen(":3000"))
}
