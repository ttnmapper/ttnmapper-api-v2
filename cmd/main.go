package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ttnmapper/ttnmapper-api-v2/internal/accountRoutes"
	"github.com/ttnmapper/ttnmapper-api-v2/internal/apiRoutes"

	"github.com/gorilla/mux"
)

func OutputUnknown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unknown request")
	fmt.Println(r)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/applications", apiRoutes.GetApplications).Methods("GET")
	router.HandleFunc("/applications/{app_id}", apiRoutes.GetApplicationDetails).Methods("GET")
	router.HandleFunc("/applications/{app_id}/devices", apiRoutes.GetDevices).Methods("GET")
	router.HandleFunc("/applications/{app_id}/devices/{dev_id}", apiRoutes.GetDeviceDetails).Methods("GET")
	router.HandleFunc("/accounts/login", accountRoutes.LoginUser).Methods("POST")
	router.HandleFunc("/accounts/checkLoginStatus", accountRoutes.CheckLoginStatus).Methods("POST")
	router.HandleFunc("/accounts/verifyToken", accountRoutes.VerifyToken).Methods("POST")
	router.HandleFunc("/", OutputUnknown).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}
