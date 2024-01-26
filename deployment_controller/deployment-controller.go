package deployment_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kube-api-comms/deployment_controller_type"
	"net/http"
	// "kube-api-comms/deployment_type"
	"github.com/gorilla/mux"
)

func CreateDeployment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body, _ := ioutil.ReadAll(request.Body)

	var deploymentCreationBody deployment_controller_type.CreateDeploymentBody

	error := json.Unmarshal(body, &deploymentCreationBody)

	if error != nil {
		fmt.Println("There was an error unmarshalling the body.")
	}

	responseMessage := deployment_controller_type.DeploymentCreationResponse{Message: "The deployment has been created succesfuly."}

	message,_ := json.Marshal(responseMessage)

	response.Write(message)
}


func DeleteDeployment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	deployment := mux.Vars(request)["deployment"]

	responseMessage := deployment_controller_type.DeploymentCreationResponse{Message: "The deployment " +deployment+ " has been deleted succesfuly."}

	message,_ := json.Marshal(responseMessage)

	response.Write(message)
}