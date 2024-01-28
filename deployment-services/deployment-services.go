package deployment_service
import ("kube-api-comms/deployment_type")

func GenerateDeploymentObject (deployment string,replicas int) (deploymentObject deployment_type.Deployment){
	deploymentObject := deployment_type.Deployment{
		ApiVersion: "apps/v1",
		Kind: "Deployment" ,
		Metadata:{
			Name:deployment,
			Labels:{
				App:"nginx"
			}
		},
		Spec:{
			Replicas:replicas,
			Selector:{
				MatchLabels:{
					App:"nginx"
				}
			},
			Template:{
				Metadata:{
					Labels:{
						App: "nginx"
					}
				},
				Spec:{
					Containers:[
						{
							Name: deployment + " - container",
							Image: "busybox:stable",
							Ports:[
								{ContainerPort:8080}
							]
						}
					]
				}
			}
		}
	}

	return
}