package deployment_controller_type

type CreateDeploymentBody struct {
	DeploymentName string `json:"deploymentname"`
	Replicas       int    `json:"replicas"`
}

type PutDeploymentBody struct {
	DeploymentName    string `json:"deploymentname"`
	Replicas          int    `json:"replicas"`
	NewDeploymentName string `json:"newdeploymentname"`
}

type DeploymentCreationResponse struct {
	Message string `json:"message"`
}

type DeploymentGetResponse struct {
	Deployments map[string]interface{} `json:"deployments"`
}
