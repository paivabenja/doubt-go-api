package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var SalesColl *mongo.Collection

func CreateSalesCollection() {
	coll := Db.Collection("sales")
	SalesColl = coll
}
