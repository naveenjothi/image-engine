package repo

import (
	"context"
	"log"
	"myapp/internal/model"
	"myapp/internal/utils/mongoutil"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user model.User) (*mongo.InsertOneResult, error) {
	collection := mongoutil.GetCollection("user")
    return collection.InsertOne(context.Background(), user)
}

func FindOneUserById(id string) (model.User, error) {
    collection := mongoutil.GetCollection("user")

    // Log the input ID
    log.Printf("Received ID: %s", id)

    // Convert string ID to ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Printf("Error converting ID to ObjectID: %v", err)
        return model.User{}, err
    }

    log.Printf("Converted ObjectID: %s", objectID.Hex())

    var user model.User

    // Find the user by ObjectID
    err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        log.Printf("Error finding user: %v", err)
        return model.User{}, err
    }

    // log.Printf("Found user: %+v", user)
    return user, nil
}

func UpdateUser(id string, user model.User) (model.User, error) {
	collection := mongoutil.GetCollection("user")
	objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return model.User{}, err
    }
   
	_, err = collection.UpdateByID(context.Background(), objectID, bson.M{
        "$set": user,
    })
	
	if err != nil {
		return model.User{}, err
	}
	var updatedUser model.User

	collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&updatedUser)
    
	return updatedUser, nil
}

func FindManyUserWithCursor() (*mongo.Cursor,error) {
	collection := mongoutil.GetCollection("user")
	cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
	return cursor, nil
}
