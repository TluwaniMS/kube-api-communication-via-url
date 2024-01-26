package deployment_controller_type 

type CreateDeploymentBody struct {
	deploymentName string `json:"deploymentname"`
	replicas       int    `json:"replicas"`
}