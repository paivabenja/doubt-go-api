package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database
var Bucket *gridfs.Bucket

func ConnectToDb(uri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
	}

	Db = client.Database("godoubt")
	Bucket, err = gridfs.NewBucket(Db)
	if err != nil {
		log.Println(err)
	}
}
