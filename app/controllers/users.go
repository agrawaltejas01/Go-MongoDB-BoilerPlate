package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/app/service"
	"github.com/gorilla/mux"

	responses "github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/lib"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {

	result, err := service.GetUsers()

	if err != nil {
		responses.InternalErrorResponse(res, err)
	}
	json.NewEncoder(res).Encode(result)
}

func GetUser(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	result, err := service.GetUser(params["id"])

	if err != nil {
		responses.InternalErrorResponse(res, err)
	}
	json.NewEncoder(res).Encode(result)
}

func AddUser(res http.ResponseWriter, req *http.Request) {
	result, err := service.AddUser(req.Body)

	if err != nil {
		responses.InternalErrorResponse(res, err)
	}
	json.NewEncoder(res).Encode(result)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	err := service.UpdateUser(params["id"])

	if err != nil {
		responses.InternalErrorResponse(res, err)
	}
	json.NewEncoder(res).Encode(true)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	err := service.DeleteUser(params["id"])

	if err != nil {
		responses.InternalErrorResponse(res, err)
	}
	json.NewEncoder(res).Encode(true)
}
