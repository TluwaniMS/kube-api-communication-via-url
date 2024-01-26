package server-set-up

func StartServer () {
	PORT := os.Getenv("PORT")

	if PORT == ""{
		PORT = "3000"
	}

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

	deploymentController.HandleFunc("/create-deployment", second_controller.PostHandler).Methods("POST")

	deploymentController.HandleFunc("/{deployment}/", second_controller.DeletHandler).Methods("DELETE")

	log.Println("Router configuration has been completed successfuly.")

	return router
}