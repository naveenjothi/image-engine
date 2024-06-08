package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var MongoClient *mongo.Client


func LoadConfig() {
    log.Println("Loading configuration...")
	err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file")
    }


    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI environment variable not set")
    }
   
	clientOptions := options.Client().ApplyURI(mongoURI)

    MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = MongoClient.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
}