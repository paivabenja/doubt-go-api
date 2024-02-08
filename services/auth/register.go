package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(password)

	res, err := database.AuthColl.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
