package middleware

import (
	"context"
	"errors"
	"rms-api/db"
	"rms-api/services/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AdminOnly(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	token := strings.Split(header, " ")[1]

	sub, err := jwt.DecodeJwtToken(token)
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "Доступ запрещён",
		})
	}

	admin := &db.Admin{}
	res := db.Client.Where("email = ?", sub).First(&admin)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "Доступ запрещён",
		})
	}

	db.Client.Where("email = ?", sub).Update("LatestIP", c.IP())
	c.Set("X-Admin-Email", admin.Email)

	ctx := context.WithValue(context.Background(), "admin", admin)
	c.SetUserContext(ctx)

	return c.Next()
}
