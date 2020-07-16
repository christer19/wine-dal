package app

import (
	"fmt"
	"net/http"

	"github.com/alvarezjulia/wine-dal/src/clients"
	controller "github.com/alvarezjulia/wine-dal/src/controllers"
	"github.com/gorilla/mux"
)

// StartApp starts the application and sets up all endpoints
// This isn't done in main.go to make it callable from testing
func StartApp() {
	router := mux.NewRouter()
	controller.MapUrls(router)

	var srv *http.Server
	srv = &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: router,
	}

	if err := clients.InitElastic(); err != nil {
		fmt.Println("error connecting to elastic cluster. Terminating service...")
		panic(err)
	}

	fmt.Println("Listening on port 8081")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("error creating http listener: " + err.Error())
		return
	}

}
