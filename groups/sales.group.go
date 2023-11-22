package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SalesGroup(group fiber.Router, client *mongo.Client) {
	database.CreateSalesCollection()

	group.Get("/", func(c *fiber.Ctx) error {
		return services.GetAllSales(c)
	})

	group.Post("/", func(c *fiber.Ctx) error {
		return services.CreateSale(c)
	})
}
