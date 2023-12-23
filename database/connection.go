package database

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/paivabenja/doubt-go-api/projectpath"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func ConnectToDb(uri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
    log.Println(1)
		panic(err)
	}

	Db = client.Database("godoubt")
	bucket, err := gridfs.NewBucket(Db)
	if err != nil {
    log.Println(2)
		panic(err)
	}

	file, err := os.Open(projectpath.ProjectRoot + "/app/archivoDePrueba.txt")
	if err != nil {
    log.Println(3)
		panic(err)
	}

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{Key: "metadata tag", Value: "first"}})

	objectId, err := bucket.UploadFromStream("a$ap.txt", io.Reader(file), uploadOpts)
	if err != nil {
		panic(err)
	}

	log.Println("me subi el archidown" + objectId.String())

}
