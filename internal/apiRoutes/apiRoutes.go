package apiRoutes

import (
	"net/http"
)

type TTNApplication struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

/**
 */
func GetApplications(w http.ResponseWriter, r *http.Request) {
}

func GetApplicationDetails(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// app_id := vars["app_id"]
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// app_id := vars["app_id"]
}

func GetDeviceDetails(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// app_id := vars["app_id"]
	// dev_id := vars["dev_id"]
}
