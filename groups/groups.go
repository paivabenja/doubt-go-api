package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/paivabenja/doubt-go-api/services"
)

func Groups(port string) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ruz!!")
	})

  app.Get("/images/:id", func (c *fiber.Ctx) error {
	 return services.GetClthImage(c)
	})

	clothesGroup := app.Group("/clothes")
	ClothesGroup(clothesGroup)

	salesGroup := app.Group("/sales")
	SalesGroup(salesGroup)

	authGroup := app.Group("/auth")
	AuthGroup(authGroup)

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
