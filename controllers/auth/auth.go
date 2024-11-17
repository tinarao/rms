package auth_controller

import (
	"errors"
	"os"
	"rms-api/db"

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

	existingAuthRequest := &db.AuthRequest{}
	result := db.Client.Where("phone = ?", data.Phone).First(&existingAuthRequest)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		authRequest := &db.AuthRequest{
			Phone: data.Phone,
		}

		db.Client.Create(authRequest)
	}

	return c.Redirect(os.Getenv("TELEGRAM_BOT_URL"))
}
