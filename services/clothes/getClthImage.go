package clothes

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetClthImage(c *fiber.Ctx) error {
	idHex := c.Params("id")

	imgId, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return err
	}

	img, err := database.Bucket.OpenDownloadStream(imgId)
	if err != nil {
		return err
	}

	fileBuffer := bytes.NewBuffer(nil)

	_, err = io.Copy(fileBuffer, img)
	if err != nil {
		return err
	}

	defer img.Close()

	return c.SendStream(fileBuffer)
}
