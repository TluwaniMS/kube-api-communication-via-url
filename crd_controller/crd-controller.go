package crd_controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"kube-api-comms/crd_controller_type"
	"kube-api-comms/http_requests"
	"kube-api-comms/kube_api_requests"
	"net/http"
)

func GetCrds(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	crd := mux.Vars(request)["customresource"]

	client := http_requests.GetClient()

	kubeApiEndPoint := kube_api_requests.GenerateCrdApi(crd)
	kubeRequest := http_requests.GenerateGetRequest(kubeApiEndPoint)

	kubeResponseBody := http_requests.MakeHttpRequest(client, kubeRequest)

	var kubeResponseObject map[string]interface{}

	error := json.Unmarshal(kubeResponseBody, &kubeResponseObject)

	if error != nil {
		fmt.Println("There was an error unmarshalling the body.")
	}

	responseMessage := crd_controller_type.CrdGetResponse{Crd: kubeResponseObject}

	message, _ := json.Marshal(responseMessage)

	response.Write(message)
}
