package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MapUrls add the endpoints of this service to the passed gorilla router
func MapUrls(router *mux.Router) {
	router.HandleFunc("/wine/{wineID}", WinesController.SearchOne).Methods(http.MethodGet)
	router.HandleFunc("/wine", WinesController.AddOne).Methods(http.MethodPost)
	router.HandleFunc("/wine/bulk", WinesController.AddMany).Methods(http.MethodPost)
}
