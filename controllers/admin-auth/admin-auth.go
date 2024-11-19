package adminauth_controller

import (
	"errors"
	"log/slog"
	"rms-api/db"
	"rms-api/services/jwt"
	"rms-api/services/validator"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

type loginDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Администратор с таким адресом электронной почты уже зарегистрирован",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to generate password hash", err)
		c.Status(fiber.StatusInternalServerError)
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
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Возникла ошибка при создании нового администратора",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Новый администратор успешно создан!",
	})

}

func Login(c *fiber.Ctx) error {
	data := &loginDTO{}
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
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Администратор не найден",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data.Password))
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Администратор не найден",
		})
	}

	token, err := jwt.GenerateJwtToken(int(admin.ID))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Возникла ошибка при авторизации",
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = *token
	cookie.Expires = time.Now().Add(time.Hour * 2)

	c.Cookie(cookie)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Вы успешно авторизовались!",
		"admin": fiber.Map{
			"firstName": admin.FirstName,
			"lastName":  admin.LastName,
			"email":     admin.Email,
		},
	})
}
