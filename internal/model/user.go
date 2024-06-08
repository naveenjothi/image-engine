package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name  string `json:"name" bson:"name"`
    Email string `json:"email" bson:"email"`
}