package model

import (
	"myapp/internal/utils/bcryptutil"
	"myapp/internal/utils/mongoutil"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name                 string             `json:"name" bson:"name"`
    Email                string             `json:"email" bson:"email"`
    Plan                 Plan             `json:"plan" bson:"plan"`
    PasswordHash         string             `json:"passwordHash" bson:"passwordHash"`
    Mobile               string             `json:"mobile" bson:"mobile"`
    IsMobileNumberVerified bool             `json:"isMobileNumberVerified" bson:"isMobileNumberVerified"`
    IsEmailVerified      bool               `json:"isEmailVerified" bson:"isEmailVerified"`
    Photo                string             `json:"photo" bson:"photo"`
	IsDeleted            bool               `json:"isDeleted" bson:"isDeleted"`
    CreatedAt            time.Time          `json:"createdAt" bson:"createdAt"`
    UpdatedAt            time.Time         `json:"updatedAt" bson:"updatedAt"`
}

type CreateUserInput struct {
    Name                 string `json:"name"`
    Email                string `json:"email"`
    Plan                 string `json:"plan"`
    Password         string `json:"passwordHash"`
    Mobile               string `json:"mobile"`
    IsMobileNumberVerified bool `json:"isMobileNumberVerified"`
    IsEmailVerified      bool `json:"isEmailVerified"`
}


func NewUser(input CreateUserInput) (User, error) {
    passwordHash, err := bcryptutil.HashPassword(input.Password)
    if err != nil {
        return User{}, err
    }

    plan := ParsePlan(input.Plan)

    now := time.Now()
    return User{
        ID:                     mongoutil.GenerateId(),
        Name:                   input.Name,
        Email:                  input.Email,
        Plan:                   plan,
        PasswordHash:           passwordHash,
        Mobile:                 input.Mobile,
        IsMobileNumberVerified: false,
        IsEmailVerified:        false,
        Photo:                  "",
        IsDeleted:              false,
        CreatedAt:              now,
        UpdatedAt:              now,
    },nil
}

type Plan int

const (
    Free Plan = iota
    Basic
    Premium
)

// String method to provide a human-readable representation of the enum values
func (p Plan) String() string {
    return [...]string{"Free", "Basic", "Premium"}[p]
}

// ParsePlan converts a string to a Plan type
func ParsePlan(plan string) Plan {
    switch plan {
    case "Basic":
        return Basic
    case "Premium":
        return Premium
    default:
        return Free
    }
}