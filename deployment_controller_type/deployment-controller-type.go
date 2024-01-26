package deployment_controller_type

type CreateDeploymentBody struct {
	deploymentName string `json:"deploymentname"`
	replicas       int    `json:"replicas"`
}

type DeploymentCreationResponse struct {
	Message string `json:"message"`
}
