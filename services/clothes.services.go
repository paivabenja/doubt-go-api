package services

import (
	"context"
	"errors"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func saveImage(c *fiber.Ctx, img_name string) (string, error) {
	img_id := uuid.NewString()
	img, err := c.FormFile(img_name)
	if err != nil {
		return "", err
	}

	c.SaveFile(img, "./public/"+img_id)
	return img_id, nil
}

func CreateClothe(coll *mongo.Collection, c *fiber.Ctx) error {
	var clothe models.Clothe
	img_back_id, err := saveImage(c, "img_back")
	if err != nil {
		return err
	}
	img_front_id, err := saveImage(c, "img_front")
	if err != nil {
		return err
	}

	c.BodyParser(&clothe)
	clothe.Img_back = img_back_id
	clothe.Img_front = img_front_id

	res, err := coll.InsertOne(context.TODO(), clothe)
	if err != nil {
		return err
	}

	if !slices.Contains(models.ClotheTypes, clothe.Type) {
		return errors.New("invalid clothe type")
	}

	return c.JSON(res)
}

func GetAllClothes(coll *mongo.Collection, c *fiber.Ctx) error {
	var clothes []models.ClotheWithId
	res, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	for res.Next(context.TODO()) {
		var clothe models.ClotheWithId
		res.Decode(&clothe)
		clothes = append(clothes, clothe)
	}

	return c.JSON(clothes)
}
