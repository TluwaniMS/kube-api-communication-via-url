package http_requests

import (
	"bytes"
	"log"
	"net/http"
)

func GetClient() *http.Client {
	client := &http.Client{}
	return client
}

func GenerateGetRequest(endPoint string) []byte {
	req, err := http.NewRequest(http.MethodGet, endpoint)

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
}

func GeneratePutRequest(endPoint string, jsonData []byte) []byte {
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
}

func GeneratePostRequest(endPoint string, jsonData []byte) []byte {
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
}
