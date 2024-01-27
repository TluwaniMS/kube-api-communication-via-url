package kube_api_requests

import ("os")

var KUBE_API = os.Getenv("KUBE_API")

if KUBE_API == ""{
	KUBE_API = "http://127.0.0.1:8001"
}

func GenerateGetDeploymentsApi(nameSpace string) string {
	deploymentApi := KUBE_API + '/api/v1/' + nameSpace + '/deployment'

	return deploymentApi
}

func GenerateCreateDeploymentApi(nameSpace string) string{
	deploymentApi := KUBE_API + '/api/v1/' + nameSpace + '/deployment'

	return deploymentApi
}

func GenerateDeleteDeploymentApi(nameSpace string) string{
	deploymentApi := KUBE_API + '/api/v1/' + nameSpace + '/deployment'

	return deploymentApi
}
