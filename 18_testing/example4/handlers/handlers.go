package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes will add the routes to http server
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON convert literal struct to json and send it to user
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Mohammad",
		Email: "k2527806@gmail.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
