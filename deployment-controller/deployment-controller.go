package deployment-controller

import (
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func CreateDeployment (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body, _ := ioutil.ReadAll(request.Body)
}

func DeleteDeployment (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	deployment := mux.Vars(request)["deployment"]
}
