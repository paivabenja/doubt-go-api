package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/paivabenja/doubt-go-api/database"
)

func Groups(port string) {
	client := database.Client
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ruz!!")
	})

	app.Static("/images", "./public")

	clothesGroup := app.Group("/clothes")
	ClothesGroup(clothesGroup, client)

	salesGroup := app.Group("/sales")
	SalesGroup(salesGroup, client)

	authGroup := app.Group("/auth")
	AuthGroup(authGroup)

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
