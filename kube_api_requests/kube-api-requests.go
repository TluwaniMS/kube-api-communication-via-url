package kube_api_requests

import "fmt"

var KUBE_API = "http://127.0.0.1:8001"

func GenerateGetDeploymentsApi(nameSpace string) string {
	deploymentApi := fmt.Sprintf("%s/api/v1/namespaces/%s/pods", KUBE_API, nameSpace)

	return deploymentApi
}

func GenerateCreateDeploymentApi(nameSpace string) string {
	deploymentApi := fmt.Sprintf("%s/api/v1/%s/deployment", KUBE_API, nameSpace)

	return deploymentApi
}

func GenerateDeleteDeploymentApi(nameSpace string) string {
	deploymentApi := fmt.Sprintf("%s/api/v1/%s/deployment", KUBE_API, nameSpace)

	return deploymentApi
}
