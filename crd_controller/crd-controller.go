package crd_controller

import (
	"net/http"
)

func GetCrds(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

}
