package auth

import (
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAdmin(c *fiber.Ctx) error {
	log.Println("adding adminnn")

	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid user Id")
	}

	var userInstance models.UserModel

	err = database.AuthColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objId}}).Decode(&userInstance)
	if err != nil {
		return errors.New("User not found")
	}

	if userInstance.IsAdmin {
		return errors.New("User is admin already")
	}

	userInstance.IsAdmin = true

	res, err := database.AuthColl.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objId}}, bson.D{{Key: "$set", Value: userInstance}})
	if err != nil {
		return err
	}

	return c.JSON(res)
}
