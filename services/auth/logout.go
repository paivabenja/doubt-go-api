package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	tokenExpiringDate := time.Now().Add(-time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  tokenExpiringDate,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	err := c.SendStatus(fiber.StatusOK)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
