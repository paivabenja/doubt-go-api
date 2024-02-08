package clothes

import (
	"io"

	"github.com/google/uuid"
	"github.com/paivabenja/doubt-go-api/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveClthImage(c *fiber.Ctx, img_name string) (string, error) {
	img_id := uuid.NewString()
	img_header, err := c.FormFile(img_name)
	if err != nil {
		return "", err
	}

	img, err := img_header.Open()
	if err != nil {
		return "", err
	}

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{
		Key:   "metadata tag",
		Value: "id: " + img_id,
	}})

	objId, err := database.Bucket.UploadFromStream(img_id, io.Reader(img), uploadOpts)
	if err != nil {
		panic(err)
	}

	return objId.Hex(), nil
}
