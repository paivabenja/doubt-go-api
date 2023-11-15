package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SalesGroup(group fiber.Router, client *mongo.Client) {
	coll := client.Database("godoubt").Collection("sales")

	group.Get("/", func(c *fiber.Ctx) error {
		return services.GetAllSales(coll, c)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		return services.CreateSale(coll, c)
	})
}
