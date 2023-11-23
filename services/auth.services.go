package services

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "dale q la vida mucho no dura"

func Register(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(password)

	res, err := database.AuthColl.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	var user models.UserModel

	res := database.AuthColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: data.Email}})
	err = res.Decode(&user)
	if err != nil {
		log.Println(err)
		err := c.SendStatus(fiber.StatusNotFound)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"message": "User not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		err := c.SendStatus(fiber.StatusBadRequest)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	tokenExpiringDate := time.Now().Add(time.Hour * 24)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.Id.Hex(),
		ExpiresAt: jwt.NewNumericDate(tokenExpiringDate),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		err := c.SendStatus(fiber.StatusInternalServerError)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"message": "Could not log in"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  tokenExpiringDate,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(user)
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		err := c.SendStatus(fiber.StatusNotFound)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"message": "unauthenticated"})
	}

	claims := token.Claims
	id, err := claims.GetIssuer()
	if err != nil {
		return err
	}

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var user models.UserModel

	res := database.AuthColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}})
	err = res.Decode(&user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	tokenExpiringDate := time.Now().Add(-time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  tokenExpiringDate,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	err := c.SendStatus(fiber.StatusOK)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var userInstance models.User
	err := c.BodyParser(&userInstance)
	if err != nil {
		log.Println("body parce")
		return err
	}

	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("objid from hex")
		return err
	}

	res, err := database.AuthColl.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objId}}, bson.D{{Key: "$set", Value: userInstance}})
	if err != nil {
		log.Println("update one")
		return err
	}

	return c.JSON(res)
}
