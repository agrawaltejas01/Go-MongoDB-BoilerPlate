package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errBody struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func sendErrorResponse(response http.ResponseWriter, status int, err error) {
	fmt.Println(err)
	body := &errBody{false, err.Error()}
	response.WriteHeader(status)
	json.NewEncoder(response).Encode(body)
}

func InternalErrorResponse(response http.ResponseWriter, err error) {
	sendErrorResponse(response, http.StatusInternalServerError, err)
}

func BadRequestResponse(response http.ResponseWriter, err error) {
	sendErrorResponse(response, http.StatusBadRequest, err)
}

func UnAuthorizedResponse(response http.ResponseWriter, err error) {
	sendErrorResponse(response, http.StatusUnauthorized, err)
}

func ValidationerrResponse(response http.ResponseWriter, err error) {
	sendErrorResponse(response, http.StatusUnprocessableEntity, err)
}
