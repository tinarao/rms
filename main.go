package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"log"
	"os"
	"rms-api/db"
	"rms-api/redis"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	app := fiber.New()
	redis.Init(redisPassword)
	db.Init()

	app.Get("/hc", func(c fiber.Ctx) error {
		return c.SendString("Ok")
	})

	log.Fatal(app.Listen(":3000"))
}
