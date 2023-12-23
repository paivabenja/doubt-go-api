package database

import "go.mongodb.org/mongo-driver/mongo"

var ClothesColl *mongo.Collection

func CreateClothesCollection() {
	coll := Db.Collection("clothes")
	ClothesColl = coll
}
