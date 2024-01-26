package server_set_up
import ("kube-api-comms/deployment_controller")

func StartServer () {
	PORT := os.Getenv("PORT")

	if PORT == ""{
		PORT = "3000"
	}

	router := ConfigureRoutes()

	server := &http.Server{
		Addr:           ":" + PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	server.ListenAndServe()
}

func ConfigureRoutes() *mux.Router {
	log.Println("Starting with Route Configuration...")
	router := mux.NewRouter()

	deploymentController := router.PathPrefix("/api/deployment-controller").Subrouter()
	
	deploymentController.HandleFunc("/create-deployment", deployment_controller.CreateDeployment).Methods("POST")
	deploymentController.HandleFunc("/{deployment}/", deployment_controller.DeleteDeployment).Methods("DELETE")

	log.Println("Router configuration has been completed successfuly.")

	return router
}