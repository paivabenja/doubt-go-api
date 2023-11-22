package groups

import (
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ClothesGroup(group fiber.Router, client *mongo.Client) {
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
