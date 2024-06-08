package model

type User struct {
    Name                 string             `json:"name" bson:"name"`
    Email                string             `json:"email" bson:"email"`
    Plan                 string             `json:"plan" bson:"plan"`
    PasswordHash         string             `json:"passwordHash" bson:"passwordHash"`
    Salt                 string             `json:"salt" bson:"salt"`
    Mobile               string             `json:"mobile" bson:"mobile"`
    IsMobileNumberVerified bool             `json:"isMobileNumberVerified" bson:"isMobileNumberVerified"`
    IsEmailVerified      bool               `json:"isEmailVerified" bson:"isEmailVerified"`
    Photo                string             `json:"photo" bson:"photo"`
    BaseModel
}
