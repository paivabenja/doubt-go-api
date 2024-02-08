package auth

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
)

func SearchUserByName(c *fiber.Ctx) error {
	type Body struct {
		Name string `json:"name"`
	}

	var users []models.UserModel

	type Filter struct {
		Name struct {
			Regex string `bson:"$regex"`
		} `bson:"name"`
	}

	body := c.Body()
	name := string(body[:])

	var filter Filter
	filter.Name.Regex = name

	res, err := database.AuthColl.Find(context.TODO(), filter)
	if err != nil {
		return errors.New("ESTO ES UN ERROR" + err.Error())
	}

	for res.Next(context.TODO()) {
		var user models.UserModel

		err := res.Decode(&user)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	return c.JSON(users)
}
