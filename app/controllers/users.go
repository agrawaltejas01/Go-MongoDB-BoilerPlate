package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/app/service"
	"github.com/gorilla/mux"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	result, err := service.GetUsers()

	if err != nil {
		fmt.Println(err)
		http.Error(res, err.Error(), 500)
	}
	json.NewEncoder(res).Encode(result)
}

func GetUser(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	result, err := service.GetUser(params["id"])

	if err != nil {
		fmt.Println(err)
		http.Error(res, err.Error(), 500)
	}
	json.NewEncoder(res).Encode(result)
}
