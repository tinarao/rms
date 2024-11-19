package restaurants_controller

import (
	"errors"
	"rms-api/db"
	"rms-api/services/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type createRestaurantDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func GetAllRestaurants(c *fiber.Ctx) error {
	restaurants := &[]db.Restaurant{}
	res := db.Client.Find(&restaurants)
	if res.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Ошибка при получении списка ресторанов",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"restaurants": restaurants,
	})
}

func CreateRestaurant(c *fiber.Ctx) error {
	data := &createRestaurantDTO{}
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

	adminEmail := c.Get("Admin-Email")
	admin := &db.Admin{}

	res := db.Client.Where("email = ?", adminEmail).First(&admin)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Ошибка аутентификации",
		})
	}

	restaurant := &db.Restaurant{
		Name:        data.Name,
		Description: data.Description,
		CreatorId:   admin.ID,
	}

	res = db.Client.Create(restaurant)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Ошибка при создании ресторана. Попробуйте позже",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Ресторан успешно создан!",
	})
}
