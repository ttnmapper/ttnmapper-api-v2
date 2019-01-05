package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TTNApplication struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var applications []TTNApplication

func GetApplications(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(applications)
}

func OutputUnknown(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Unknown request")
    fmt.Println(r)
}

// our main function
func main() {

	// Add some demo data
	applications = append(applications, TTNApplication{ID: "001", Name: "demo-gateway-1"})
	applications = append(applications, TTNApplication{ID: "002", Name: "demo-gateway-2"})
	applications = append(applications, TTNApplication{ID: "003", Name: "demo-gateway-3"})

	router := mux.NewRouter()
	router.HandleFunc("/applications", GetApplications).Methods("GET")
	router.HandleFunc("/", OutputUnknown).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
