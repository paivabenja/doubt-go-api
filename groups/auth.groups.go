package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/services"
)

func AuthGroup(group fiber.Router) {
	database.CreateAuthCollection()

	group.Post("/register", services.Register)
	group.Post("/login", services.Login)
	group.Get("/user", services.User)
	group.Get("/logout", services.Logout)
	group.Put("/:id", services.UpdateUser)
}
