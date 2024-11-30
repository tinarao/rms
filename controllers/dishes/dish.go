package dishes_controller

import (
	"errors"
	"log/slog"
	"rms-api/db"
	"rms-api/services/validator"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type createDishDTO struct {
	Name         string  `json:"name" validate:"required,min=4,max=32"`
	Description  *string `json:"description"`
	RestaurantId uint    `json:"restaurantId" validate:"required"`
}

func Create(c *fiber.Ctx) error {
	adminctx := c.UserContext().Value("admin")
	admin, ok := adminctx.(*db.Admin)
	if !ok {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Ошибка аутентификации",
		})
	}

	data := &createDishDTO{}
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

	pNameSlug := slug.Make(data.Name)

	dup := &db.Dish{}
	res := db.Client.Where("name = ?", data.Name).Or("slug = ?", pNameSlug).First(&dup)
	if res.Error == nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Продукт с похожим названием уже существует",
		})
	}

	desc := ""
	if data.Description != nil {
		desc = *data.Description
	}

	dish := &db.Dish{
		Name:         data.Name,
		Description:  &desc,
		RestaurantID: data.RestaurantId,
		CreatorId:    admin.ID,
		CreatedAt:    time.Now(),
		Slug:         pNameSlug,
	}
	res = db.Client.Create(dish)
	if res.Error != nil {
		slog.Error(res.Error.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Возникла ошибка при создании продукта",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Продукт создан!",
	})
}

func GetPrivateDataBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	dish := &db.Dish{}
	res := db.Client.Where("slug = ?", slug).Preload("Orders").Preload("Restaurant").First(&dish)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "Блюдо не найдено",
			})
		}
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"dish": dish,
	})
}
