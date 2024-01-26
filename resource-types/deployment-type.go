package deployment-type 

type Deployment struct {
	apiVersion string `json:"apiversion"`
	kind       string `json:"kind"`
	metadata   Metadata `json:"metadata"`
	spec       Spec `json:"spec"`
}

type Metadata struct {
	name   string `json:"name"`
	labels map[string]string `json:"labels"`
}

type Spec struct {
	replicas int `json:"replicas"`
	selector Selector `json:"selector"`
	template Template `json:"template"`
}

type Selector struct {
	matchLabels map[string]string `json:"matchlabels"`
}

type Template struct {
	metadata TemplateMetadata `json:"metadata"`
	spec     TemplateSpec `json:"spec"`
}

type TemplateMetadata struct {
	labels map[string]string `json:"labels"`
}

type TemplateSpec struct {
	containers []Container `json:"containers"`
}

type Container struct {
	name  string `json:"name"`
	image string `json:"image"`
	ports []ContainerPort `json:"ports"`
}

type ContainerPort struct {
	containerPort int `json:"containerport"`
}