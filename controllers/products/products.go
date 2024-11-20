package products_controller

import (
	"github.com/gofiber/fiber/v2"
)

type createProductDto struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Tags        []string `json:"tags"`
}

func CreateProduct(c *fiber.Ctx) error {
	// adminctx := c.UserContext().Value("admin")
	// admin, ok := adminctx.(*db.Admin)
	// if !ok {
	// 	c.Status(fiber.StatusInternalServerError)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Ошибка аутентификации",
	// 	})
	// }

	// data := &createProductDto{}
	// if err := c.BodyParser(&data); err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Bad request",
	// 	})
	// }

	// err := validator.Validate.Struct(data)
	// if err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Форма заполнена некорректно",
	// 	})
	// }

	// pNameSlug := slug.Make(data.Name)

	// dup := &db.Product{}
	// res := db.Client.Where("name = ?", data.Name).Or("slug = ?", pNameSlug).First(&dup)
	// if res.Error != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "Продукт с похожим названием уже существует",
	// 	})
	// }
}
