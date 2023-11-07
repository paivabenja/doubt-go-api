package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/paivabenja/doubt-go-api/groups"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getEnvs() (string, string) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	mongo_uri := os.Getenv("MONGODB_URI")

	return port, mongo_uri
}

func connectToDb(uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	port, mongo_uri := getEnvs()
	client := connectToDb(mongo_uri)

	groups.Groups(port, client)
}
