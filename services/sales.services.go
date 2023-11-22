package services

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateSale(c *fiber.Ctx) error {
	var sale models.Sale
	err := c.BodyParser(&sale)
	if err != nil {
		return err
	}

	date := time.Now().Format("2006-1-2 15:4:5")
	sale.Date = date

	res, err := database.SalesColl.InsertOne(context.TODO(), sale)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func GetAllSales(c *fiber.Ctx) error {
	var sales []models.Sale
	res, err := database.SalesColl.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	for res.Next(context.TODO()) {
		var sale models.Sale
		err := res.Decode(&sale)
		if err != nil {
			return err
		}
		sales = append(sales, sale)
	}

	return c.JSON(sales)
}
