package restaurants_controller

import (
	"errors"
	"rms-api/db"
	"rms-api/services/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
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
	adminctx := c.UserContext().Value("admin")
	admin, ok := adminctx.(*db.Admin)
	if !ok {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Ошибка аутентификации",
		})
	}

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

	rNameSlug := slug.Make(data.Name)
	restaurant := &db.Restaurant{
		Name:        data.Name,
		Description: data.Description,
		CreatorId:   admin.ID,
		Slug:        rNameSlug,
	}

	res := db.Client.Create(restaurant)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Ресторан с таким названием уже существует",
			})
		}

		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Ошибка при создании ресторана. Попробуйте позже",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Ресторан успешно создан!",
		"slug":    rNameSlug,
	})
}

func GetRestaurantBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	restaurant := &db.Restaurant{}

	res := db.Client.Where("slug = ?", slug).First(&restaurant)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Ресторан не найден",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"restaurant": restaurant,
	})
}
