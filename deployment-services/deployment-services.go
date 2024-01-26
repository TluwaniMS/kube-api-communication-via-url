package deployment_service
import ("kube-api-comms/deployment_type")

func GenerateDeploymentObject (deployment string,replicas int,image string) {
	deploymentObject := deployment_type.Deployment{
		apiVersion: "apps/v1",
		kind: "Deployment" ,
		metadata:{
			name:deployment,
			labels:{
				app:"nginx"
			}
		},
		spec:{
			replicas:replicas,
			selector:{
				matchLabels:{
					app:"nginx"
				}
			},
			template:{
				metadata:{
					labels:{
						app: "nginx"
					}
				},
				spec:{
					containers:[
						{
							name: deployment + " - container",
							image: image,
							ports:[
								{containerPort:8080}
							]
						}
					]
				}
			}
		}
	}

	
}