# K8s communication via Http requests.

## About Repository:

This repository serves as a foundational project I've established to explore interaction with the Kube-api via HTTP requests from a Golang server, enabling operations such as reading, creating, updating, and deleting Kubernetes resources.

## Software(s) required:

* [Golang](https://go.dev/doc/install)
* [Minikube](https://minikube.sigs.k8s.io/docs/start/)

## Project Set-up:

`Note:`

Make sure you have a running Minikube instance.

* #### Step 1: 

Start a proxy server to enable direct access to the Kubernetes API server from your local machine.

```
kubectl proxy --port=8081
```

* #### Step 2: 

Navigate to the root directory of the project and execute:

```
go run .
```

This will initialize the REST API server.
