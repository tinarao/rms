package auth_controller

import (
	"context"
	"errors"
	"rms-api/db"
	"rms-api/redis"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthDTO struct {
	Phone string `json:"phone"`
}

func Auth(c *fiber.Ctx) error {
	data := &AuthDTO{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := &db.User{}
	result := db.Client.Where("phone = ?", data.Phone).First(&user)
	if result.Error == nil {
		sessionId := nanoid.New()
		redis.Client.Set(context.Background(), sessionId, user.ID, 0)

		db.Client.Where("id = ?", user.ID).Update("sessionId", sessionId)

		cookie := new(fiber.Cookie)
		cookie.Name = "session_id"
		cookie.Value = sessionId
		cookie.Expires = time.Now().Add(24 * time.Hour)

		c.Cookie(cookie)
		c.Status(200)

		return c.JSON(fiber.Map{
			"phone": user.Phone,
		})
	}

	existingAuthRequest := &db.AuthRequest{}
	result = db.Client.Where("phone = ?", data.Phone).First(&existingAuthRequest)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		authRequest := &db.AuthRequest{
			Phone: data.Phone,
		}

		db.Client.Create(authRequest)
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"message":       "Вход выполнен!",
		"auth_redirect": "false",
	})
}
