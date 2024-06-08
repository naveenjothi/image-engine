package mongoutil

import (
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"myapp/internal/config"
)


func GetCollection(collectionName string) *mongo.Collection {
    return config.MongoClient.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
}

func GenerateId() primitive.ObjectID{
	return primitive.NewObjectID()
}