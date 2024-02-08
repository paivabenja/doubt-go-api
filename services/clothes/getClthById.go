package clothes

import (
	"context"

	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func stringToObjectId(objectId string) (primitive.ObjectID, error) {
	res, err := primitive.ObjectIDFromHex(objectId)
	return res, err
}

func GetClotheById(c *fiber.Ctx) error {
	var clothe models.ClotheWithId

	clotheId, err := stringToObjectId(c.Params("id"))
	if err != nil {
		return err
	}

	err = database.ClothesColl.FindOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: clotheId}},
	).Decode(&clothe)
	if err != nil {
		return err
	}

	return c.JSON(clothe)
}
