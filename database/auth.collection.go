package database

import "go.mongodb.org/mongo-driver/mongo"

var AuthColl *mongo.Collection

func CreateAuthCollection() {
	coll := Db.Collection("userAuth")
	AuthColl = coll
}
