package clothes

import (
	"context"

	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllClothes(c *fiber.Ctx) error {
	var clothes []models.ClotheWithId
	res, err := database.ClothesColl.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	for res.Next(context.TODO()) {
		var clothe models.ClotheWithId
		err := res.Decode(&clothe)
		if err != nil {
			return nil
		}
		clothes = append(clothes, clothe)
	}

	return c.JSON(clothes)
}
