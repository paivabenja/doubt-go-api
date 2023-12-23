package services

import (
	"context"
	"errors"
	"slices"

	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func saveImage(c *fiber.Ctx, img_name string) (string, error) {
	img_id := uuid.NewString()
	img, err := c.FormFile(img_name)
	if err != nil {
		return "", err
	}

	err = c.SaveFile(img, "./public/"+img_id)
	if err != nil {
		return "", err
	}
	return img_id, nil
}

func CreateClothe(c *fiber.Ctx) error {
	var clothe models.Clothe
	img_back_id, err := saveImage(c, "img_back")
	if err != nil {
		return err
	}

	img_front_id, err := saveImage(c, "img_front")
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

func GetClotheById(c *fiber.Ctx) error {
	var clothe models.ClotheWithId

	clotheId, err := stringToObjectId(c.Params("id"))
	if err != nil {
		return err
	}

	err = database.ClothesColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: clotheId}}).Decode(&clothe)
	if err != nil {
		return err
	}

	return c.JSON(clothe)
}

func stringToObjectId(objectId string) (primitive.ObjectID, error) {
	res, err := primitive.ObjectIDFromHex(objectId)
	return res, err
}
