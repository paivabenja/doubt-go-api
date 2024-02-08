package auth

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(c *fiber.Ctx) error {
	var userInstance models.User
	err := c.BodyParser(&userInstance)
	if err != nil {
		log.Println("body parce")
		return err
	}

	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("objid from hex")
		return err
	}

	res, err := database.AuthColl.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objId}}, bson.D{{Key: "$set", Value: userInstance}})
	if err != nil {
		log.Println("update one")
		return err
	}

	return c.JSON(res)
}
