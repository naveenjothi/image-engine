package service

import (
	"myapp/internal/app/repo"
	"myapp/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(id string, user model.User) error {
    return repo.UpdateUser(id, user)
}

func CreateUser(user model.User) (*mongo.InsertOneResult, error){
    return repo.CreateUser(user)
}

func FindOneUserById(id string) (model.User, error){
    return repo.FindOneUserById(id)
}