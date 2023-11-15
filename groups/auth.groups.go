package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/services"
)

func AuthGroup(group fiber.Router) {
	group.Post("/", func(c *fiber.Ctx) error {
		return services.Register(c)
	})
}
