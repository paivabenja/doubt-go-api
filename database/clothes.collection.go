package database

import "go.mongodb.org/mongo-driver/mongo"

var ClothesColl *mongo.Collection

func CreateClothesCollection() {
	coll := Client.Database("godoubt").Collection("clothes")
	ClothesColl = coll
}
