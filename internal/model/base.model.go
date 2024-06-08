package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct{
    ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IsDeleted            bool               `json:"isDeleted" bson:"isDeleted"`
    CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
    UpdatedAt            time.Time          `json:"updatedAt" bson:"updatedAt"`
}