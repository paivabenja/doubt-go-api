package auth

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.UserModel

	res := database.AuthColl.FindOne(
		context.TODO(),
		bson.D{{Key: "email", Value: data.Email}},
	)

	err = res.Decode(&user)
	if err != nil {
		log.Println(err)
		err := c.SendStatus(fiber.StatusNotFound)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"message": "User not found"})
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(data.Password),
	)

	if err != nil {
		err := c.SendStatus(fiber.StatusBadRequest)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	tokenExpiringDate := time.Now().Add(time.Hour * 24 * 7) // 1 week

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.Id.Hex(),
		ExpiresAt: jwt.NewNumericDate(tokenExpiringDate),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		err := c.SendStatus(fiber.StatusInternalServerError)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"message": "Could not log in"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  tokenExpiringDate,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(user)
}
