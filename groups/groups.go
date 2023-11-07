package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

func Groups(port string, client *mongo.Client) {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ruth!!!")
	})

	imagesGroup := app.Group("/images")
	ImagesGroup(imagesGroup, client, app)

	clothesGroup := app.Group("/clothes")
	ClothesGroup(clothesGroup, client)

	app.Listen(":" + port)
}
