package handler

import (
	"encoding/json"
	"log"
	"myapp/internal/app/service"
	"myapp/internal/model"
	"myapp/internal/utils/jsonutil"
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
    var input model.CreateUserInput
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        response.JSON(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    user,err := model.NewUser(input)
    if err != nil {
        response.JSON(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    _,err = service.CreateUser(user)
    if err != nil {
        response.JSON(w, http.StatusInternalServerError, err.Error())
        return
    }
    response.JSON(w, http.StatusCreated, user)
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var input model.UpdateUserInput
    params := mux.Vars(r)
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        response.JSON(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    user, err := service.FindOneUserById(params["id"])
    if err != nil {
        response.JSON(w, http.StatusNotFound, err.Error())
        return
    }

    jsonutil.MergeStructs(&user,&input)
    
    err = service.UpdateUser(params["id"],user)
    if err != nil {
        response.JSON(w, http.StatusInternalServerError, err.Error())
        return
    }

    log.Printf("Updated user %v",user)

    response.JSON(w, http.StatusCreated, user)
}