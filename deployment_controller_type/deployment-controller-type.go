package deployment_controller_type

type CreateDeploymentBody struct {
	deploymentName string `json:"deploymentname"`
	replicas       int    `json:"replicas"`
}

type DeploymentCreationResponse struct {
	Message string `json:"message"`
}

type DeploymentGetResponse struct {
	Deployments map[string]interface{} `json:"deployments"`
}
