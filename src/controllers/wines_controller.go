package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/alvarezjulia/wine-dal/src/domain"
	"github.com/gorilla/mux"
)

// WinesController is an interface that wraps the winesController struct type.
// Purpose: Can be mocked for testing
var WinesController WinesControllerInterface = &winesController{}

type winesController struct{}

// WinesControllerInterface is used for mocking purposes
type WinesControllerInterface interface {
	SearchOne(w http.ResponseWriter, r *http.Request)
	AddOne(w http.ResponseWriter, r *http.Request)
	AddMany(w http.ResponseWriter, r *http.Request)
}

func (e *winesController) SearchOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wineID := vars["wineID"]

	fmt.Println("received request from host: " + r.Host + " RequestURI: " + r.RequestURI)
	fmt.Println(wineID)

	wineResult, err := domain.EntryDao.GetDoc(r.Context(), wineID)
	if err != nil {
		respondJSON("", w, err)
		panic(err)
	}
	var wineResultBytes []byte
	wineResultBytes, err = json.Marshal(wineResult)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(wineResultBytes)
}

func (e *winesController) AddOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request from host: " + r.Host + " RequestURI: " + r.RequestURI)

	// enforce a maximum read of 1MB from the json body
	// this limit wrong data insertion
	jsonBody, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println("request body has no json object")
		return
	}

	var wine domain.Wine

	err = json.Unmarshal(jsonBody, &wine)
	if err != nil {
		fmt.Println("bad request: " + err.Error())
		return
	}
	var newID string
	newID, err = domain.EntryDao.CreateNewDoc(r.Context(), wine)
	if err != nil {
		respondJSON("", w, err)
		panic(err)
	}
	respondJSON("wine indexed with ID: "+newID, w, nil)
}

func (e *winesController) AddMany(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request from host: " + r.Host + " RequestURI: " + r.RequestURI)

	// enforce a maximum read of 10MB from the json body
	jsonBody, err := ioutil.ReadAll(io.LimitReader(r.Body, 10485760))
	if err != nil {
		fmt.Println("request body has no json object")
		return
	}

	var wines domain.WineList

	err = json.Unmarshal(jsonBody, &wines)
	if err != nil {
		fmt.Println("bad request: " + err.Error())
		return
	}
	err = domain.EntryDao.CreateManyDocs(r.Context(), wines)
	if err != nil {
		respondJSON("", w, err)
		panic(err)
	}
	respondJSON("sucessful bulk upload", w, nil)
}

func respondJSON(msg string, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": ` + msg + `"}`))
	}
	if err != nil {
		w.Write([]byte(`{"message": "following error in request: ` + err.Error() + `"}`))
	}
}
