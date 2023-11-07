package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func ImagesGroup(group fiber.Router, client *mongo.Client, app *fiber.App) {
	app.Static("/images", "./public")

	coll := client.Database("godoubt").Collection("images")

	group.Post("/", func(c *fiber.Ctx) error {
		return services.PostImage(c, coll)
	})
}
