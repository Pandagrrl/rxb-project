package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", WelcomeHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(WelcomeResponse{Message: "Welcome to Mockbuster!"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

type WelcomeResponse struct {
	Message string `json:"message"`
}
