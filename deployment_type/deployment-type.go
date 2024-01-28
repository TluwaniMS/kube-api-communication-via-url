package deployment_type 

type Deployment struct {
	ApiVersion string `json:"apiversion"`
	Kind       string `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec `json:"spec"`
}

type Metadata struct {
	Name   string `json:"name"`
	Labels map[string]string `json:"labels"`
}

type Spec struct {
	Replicas int `json:"replicas"`
	Selector Selector `json:"selector"`
	Template Template `json:"template"`
}

type Selector struct {
	MatchLabels map[string]string `json:"matchlabels"`
}

type Template struct {
	Metadata TemplateMetadata `json:"metadata"`
	Spec     TemplateSpec `json:"spec"`
}

type TemplateMetadata struct {
	Labels map[string]string `json:"labels"`
}

type TemplateSpec struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Ports []ContainerPort `json:"ports"`
}

type ContainerPort struct {
	ContainerPort int `json:"containerport"`
}