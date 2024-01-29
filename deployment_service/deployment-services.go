package deployment_service

import (
	"kube-api-comms/deployment_type"
)

func GenerateDeploymentObject(deployment string, replicas int) *deployment_type.Deployment {
	deploymentObject := &deployment_type.Deployment{
		ApiVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: deployment_type.Metadata{
			Name: deployment,
			Labels: deployment_type.Labels{
				App: "nginx",
			},
		},
		Spec: deployment_type.Spec{
			Replicas: replicas,
			Selector: deployment_type.Selector{
				MatchLabels: deployment_type.Labels{
					App: "nginx",
				},
			},
			Template: deployment_type.Template{
				Metadata: deployment_type.Metadata{
					Labels: deployment_type.Labels{
						App: "nginx",
					},
				},
				Spec: deployment_type.ContainerSpec{
					Containers: []deployment_type.Container{
						{
							Name:  deployment + "-container",
							Image: "busybox:stable",
							Ports: []deployment_type.Port{
								{ContainerPort: 8080},
							},
						},
					},
				},
			},
		},
	}

	return deploymentObject
}
