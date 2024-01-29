package deployment_type

type Deployment struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

type Metadata struct {
	Name   string `json:"name"`
	Labels Labels `json:"labels"`
}

type Labels struct {
	App string `json:"app"`
}

type Spec struct {
	Replicas int      `json:"replicas"`
	Selector Selector `json:"selector"`
	Template Template `json:"template"`
}

type Selector struct {
	MatchLabels Labels `json:"matchLabels"`
}

type Template struct {
	Metadata Metadata      `json:"metadata"`
	Spec     ContainerSpec `json:"spec"`
}

type ContainerSpec struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Ports []Port `json:"ports"`
}

type Port struct {
	ContainerPort int `json:"containerPort"`
}
