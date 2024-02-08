package clothes

import (
	"context"
	"errors"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
)

func CreateClothe(c *fiber.Ctx) error {
	var clothe models.Clothe
	img_back_id, err := SaveClthImage(c, "img_back")
	if err != nil {
		return err
	}

	img_front_id, err := SaveClthImage(c, "img_front")
	if err != nil {
		return err
	}

	// Parse the json body of the HTTP request
	err = c.BodyParser(&clothe)
	if err != nil {
		return err
	}
	clothe.Img_back = img_back_id
	clothe.Img_front = img_front_id

	// Insert clothe into database
	res, err := database.ClothesColl.InsertOne(context.TODO(), clothe)
	if err != nil {
		return err
	}

	// Validate clothe type
	if !slices.Contains(models.ClotheTypes, clothe.Type) {
		return errors.New("invalid clothe type")
	}

	return c.JSON(res.InsertedID)
}
