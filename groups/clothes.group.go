package groups

import (
	"github.com/paivabenja/doubt-go-api/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ClothesGroup(group fiber.Router, client *mongo.Client) {
	coll := client.Database("godoubt").Collection("clothes")

	group.Get("/", func(c *fiber.Ctx) error {
		return services.GetAllClothes(coll, c)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		return services.CreateClothe(coll, c)
	})

	group.Get("/:id", func(c *fiber.Ctx) error {
		return services.GetClotheById(coll, c)
	})
}
