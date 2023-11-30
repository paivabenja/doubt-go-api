package groups

import (
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/services"

	"github.com/gofiber/fiber/v2"
)

func ClothesGroup(group fiber.Router) {
	database.CreateClothesCollection()

	group.Get("/", func(c *fiber.Ctx) error {
		return services.GetAllClothes(c)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		return services.CreateClothe(c)
	})

	group.Get("/:id", func(c *fiber.Ctx) error {
		return services.GetClotheById(c)
	})
}
