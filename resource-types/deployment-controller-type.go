package deployment-controller-type 

type CreateDeploymentBody struct {
	deploymentName string `json:"deploymentname"`
	replicas       int    `json:"replicas"`
}