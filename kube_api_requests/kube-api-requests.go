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

func GenerateDeploymentApi(nameSpace string, deploymentName string) string {
	deploymentApi := fmt.Sprintf("%s/apis/apps/v1/namespaces/%s/deployments/%s", KUBE_API, nameSpace, deploymentName)

	return deploymentApi
}

func GenerateCrdApi(customResourceDefinition string) string {
	crdApi := fmt.Sprintf("%s/apis/apiextensions.k8s.io/v1/customresourcedefinitions/%s", KUBE_API, customResourceDefinition)

	return crdApi
}
