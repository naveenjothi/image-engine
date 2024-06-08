package handler

import (
	"encoding/json"
	"myapp/internal/app/service"
	"myapp/internal/model"
	"myapp/internal/utils/mongoutil"
	"myapp/pkg/response"
	"net/http"

	"github.com/gorilla/mux"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    user, err := service.FindOneUserById(params["id"])
    if err != nil {
        response.JSON(w, http.StatusNotFound, err.Error())
        return
    }
    response.JSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    user.ID = mongoutil.GenerateId()
    if err != nil {
        response.JSON(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    _,err = service.CreateUser(user)
    if err != nil {
        response.JSON(w, http.StatusInternalServerError, err.Error())
        return
    }
    response.JSON(w, http.StatusCreated, user)
}
