package database

import "go.mongodb.org/mongo-driver/mongo"

var SalesColl *mongo.Collection

func CreateSalesCollection() {
	coll := Client.Database("godoubt").Collection("sales")
	SalesColl = coll
}
