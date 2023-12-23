package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/paivabenja/doubt-go-api/database"
	"github.com/paivabenja/doubt-go-api/groups"
)

// TODO: save images in db

func getEnvs() (string, string) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	mongo_uri := os.Getenv("MONGODB_URI")

	return port, mongo_uri
}

func main() {
	port, mongo_uri := getEnvs()
	database.ConnectToDb(mongo_uri)
	groups.Groups(port)
}
