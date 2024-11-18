package adminauth_controller

import (
	"log/slog"
	"rms-api/db"
	"rms-api/services/validator"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

/*
	Admin authorization system uses JWT strategy
*/

type registerDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	data := &registerDTO{}
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	err := validator.Validate.Struct(data)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Форма заполнена некорректно",
		})
	}

	admin := &db.Admin{}
	result := db.Client.Where("email = ?", data.Email).First(&admin)
	if result.Error == nil {
		return c.JSON(fiber.Map{
			"message": "Администратор с таким адресом электронной почты уже зарегистрирован",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to generate password hash", err)
		return c.JSON(fiber.Map{
			"message": "Ошибка при создании токенов",
		})
	}

	createdAdmin := &db.Admin{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  string(hash),
	}

	result = db.Client.Create(createdAdmin)
	if result.Error != nil {
		slog.Error("Возникла ошибка при создании нового администратора", result.Error)
		return c.JSON(fiber.Map{
			"message": "Возникла ошибка при создании нового администратора",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Новый администратор успешно создан!",
	})

}
