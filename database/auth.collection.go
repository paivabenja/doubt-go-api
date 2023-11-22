package database

import "go.mongodb.org/mongo-driver/mongo"

var AuthColl *mongo.Collection

func CreateAuthCollection() {
	coll := Client.Database("godoubt").Collection("userAuth")
	AuthColl = coll
}
