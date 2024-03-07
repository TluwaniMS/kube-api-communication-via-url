package server_set_up

import (
	"kube-api-comms/crd_controller"
	"kube-api-comms/deployment_controller"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	router := ConfigureRoutes()

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      loggedRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

func ConfigureRoutes() *mux.Router {
	log.Println("Starting with Route Configuration...")
	router := mux.NewRouter()

	deploymentController := router.PathPrefix("/api/deployment-controller").Subrouter()
	crdController := router.PathPrefix("/api/crd-controller").Subrouter()

	deploymentController.HandleFunc("/create-deployment", deployment_controller.CreateDeployment).Methods("POST")
	deploymentController.HandleFunc("/update-deployment", deployment_controller.PutDeployment).Methods("PUT")
	deploymentController.HandleFunc("/{deployment}", deployment_controller.DeleteDeployment).Methods("DELETE")
	deploymentController.HandleFunc("/{namespace}", deployment_controller.GetDeployments).Methods("GET")
	deploymentController.HandleFunc("/{namespace}/{deployment}", deployment_controller.GetDeployment).Methods("GET")

	crdController.HandleFunc("/{customresource}", crd_controller.GetCrds).Methods("GET")

	log.Println("Router configuration has been completed successfuly.")

	return router
}
