package groups

import (
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/services/clothes"

	"github.com/gofiber/fiber/v2"
)

func ClothesGroup(group fiber.Router) {
	database.CreateClothesCollection()

	group.Get("/", func(c *fiber.Ctx) error {
		return clothes.GetAllClothes(c)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		return clothes.CreateClothe(c)
	})

	group.Get("/:id", func(c *fiber.Ctx) error {
		return clothes.GetClotheById(c)
	})
}
