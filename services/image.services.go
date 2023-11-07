package services

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostImage(c *fiber.Ctx, coll *mongo.Collection) error {
	file, err := c.FormFile("img-front")
	if err != nil {
		return err
	}

	imgId := uuid.NewString()

	_, err = coll.InsertOne(context.TODO(), bson.D{{Key: "imageId", Value: imgId}})
	if err != nil {
		return err
	}

	return c.SaveFile(file, "./public/"+imgId)
}
