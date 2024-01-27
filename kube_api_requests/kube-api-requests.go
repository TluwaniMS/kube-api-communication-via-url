package kube_api_requests

import "fmt"

var KUBE_API = "http://127.0.0.1:8081"

func GenerateGetDeploymentsApi(nameSpace string) string {
	deploymentApi := fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/deployments", KUBE_API, nameSpace)

	return deploymentApi
}

func GenerateCreateDeploymentApi(nameSpace string) string {
	deploymentApi := fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/deployments", KUBE_API, nameSpace)

	return deploymentApi
}

func GenerateDeleteDeploymentApi(nameSpace string) string {
	deploymentApi := fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/deployments", KUBE_API, nameSpace)

	return deploymentApi
}
