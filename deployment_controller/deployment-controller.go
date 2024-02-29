package deployment_controller

import (
	"encoding/json"
	"fmt"
	"io"
	"kube-api-comms/deployment_controller_type"
	"kube-api-comms/deployment_service"
	"kube-api-comms/http_requests"
	"kube-api-comms/kube_api_requests"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateDeployment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body, _ := io.ReadAll(request.Body)

	var deploymentCreationBody deployment_controller_type.CreateDeploymentBody

	error := json.Unmarshal(body, &deploymentCreationBody)

	if error != nil {
		fmt.Println("There was an error unmarshalling the body.")
	}

	deploymentObject := deployment_service.GenerateDeploymentObject(deploymentCreationBody.DeploymentName, deploymentCreationBody.Replicas)

	client := http_requests.GetClient()

	deploymentJsonObject, error := json.Marshal(deploymentObject)

	if error != nil {
		fmt.Println("There was an error marshalling the body.")
	}

	kubeApiEndPoint := kube_api_requests.GenerateGetDeploymentsApi("default")
	kubeRequest := http_requests.GeneratePostRequest(kubeApiEndPoint, deploymentJsonObject)

	kubeResponseBody := http_requests.MakeHttpRequest(client, kubeRequest)

	var kubeResponseObject map[string]interface{}

	error = json.Unmarshal(kubeResponseBody, &kubeResponseObject)

	responseMessage := deployment_controller_type.DeploymentCreationResponse{Message: "The deployment has been created succesfuly."}

	message, _ := json.Marshal(responseMessage)

	response.Write(message)
}

func PutDeployment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	body, _ := io.ReadAll(request.Body)

	var deploymentPutBody deployment_controller_type.PutDeploymentBody

	error := json.Unmarshal(body, &deploymentPutBody)

	if error != nil {
		fmt.Println("There was an error unmarshalling the body.")
	}

	deploymentObject := deployment_service.GenerateDeploymentObject(deploymentPutBody.DeploymentName, deploymentPutBody.Replicas)

	client := http_requests.GetClient()

	deploymentJsonObject, error := json.Marshal(deploymentObject)

	if error != nil {
		fmt.Println("There was an error marshalling the body.")
	}

	kubeApiEndPoint := kube_api_requests.GenerateDeploymentApi("default", deploymentPutBody.DeploymentName)
	kubeRequest := http_requests.GeneratePutRequest(kubeApiEndPoint, deploymentJsonObject)

	kubeResponseBody := http_requests.MakeHttpRequest(client, kubeRequest)

	var kubeResponseObject map[string]interface{}

	error = json.Unmarshal(kubeResponseBody, &kubeResponseObject)

	if error != nil {
		fmt.Println("There was an error unmarshalling the response body.")
	}

	fmt.Println(kubeResponseObject)

	responseMessage := deployment_controller_type.DeploymentCreationResponse{Message: "The deployment has been updated succesfuly."}

	message, _ := json.Marshal(responseMessage)

	response.Write(message)
}

func DeleteDeployment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	deployment := mux.Vars(request)["deployment"]

	responseMessage := deployment_controller_type.DeploymentCreationResponse{Message: "The deployment " + deployment + " has been deleted succesfuly."}

	client := http_requests.GetClient()

	kubeApiEndPoint := kube_api_requests.GenerateDeploymentApi("default", deployment)
	kubeRequest := http_requests.GenerateDeleteRequest(kubeApiEndPoint)

	kubeResponseBody := http_requests.MakeHttpRequest(client, kubeRequest)

	var kubeResponseObject map[string]interface{}

	error := json.Unmarshal(kubeResponseBody, &kubeResponseObject)

	if error != nil {
		fmt.Println("There was an error unmarshalling the response body.")
	}

	message, _ := json.Marshal(responseMessage)

	response.Write(message)
}

func GetDeployments(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	namespace := mux.Vars(request)["namespace"]

	kubeApiEndPoint := kube_api_requests.GenerateGetDeploymentsApi(namespace)
	kubeRequest := http_requests.GenerateGetRequest(kubeApiEndPoint)

	client := http_requests.GetClient()

	kubeResponseBody := http_requests.MakeHttpRequest(client, kubeRequest)

	var kubeResponseObject map[string]interface{}

	error := json.Unmarshal(kubeResponseBody, &kubeResponseObject)

	if error != nil {
		fmt.Println("There was an error unmarshalling the response body.")
	}

	responseMessage := deployment_controller_type.DeploymentGetResponse{Deployments: kubeResponseObject}

	message, _ := json.Marshal(responseMessage)

	response.Write(message)
}
