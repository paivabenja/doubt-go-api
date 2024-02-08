package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/services/auth"
)

func AuthGroup(group fiber.Router) {
	database.CreateAuthCollection()

	group.Post("/register", auth.Register)
	group.Post("/login", auth.Login)
	group.Get("/user", auth.GetUser)
	group.Get("/logout", auth.Logout)
	group.Put("/:id", auth.UpdateUser)
	group.Put("/admin/:id", auth.AddAdmin)
	group.Get("/name", auth.SearchUserByName)
}
