package http_requests

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func GetClient() *http.Client {
	client := &http.Client{}
	return client
}

func GenerateGetRequest(endPoint string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, endPoint, nil)

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	return req
}

func GeneratePutRequest(endPoint string, jsonData []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPut, endPoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	return req
}

func GeneratePostRequest(endPoint string, jsonData []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPost, endPoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	return req
}

func MakeHttpRequest(client *http.Client, request *http.Request) []byte {
	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}
