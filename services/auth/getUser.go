package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		err := c.SendStatus(fiber.StatusNotFound)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"message": "unauthenticated"})
	}

	claims := token.Claims
	id, err := claims.GetIssuer()
	if err != nil {
		return err
	}

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var user models.UserModel

	res := database.AuthColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}})
	err = res.Decode(&user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
