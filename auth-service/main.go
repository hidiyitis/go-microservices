package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)
	if req.Username == "admin" && req.Password == "password" {
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "token": "dummy-jwt-token"})
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"message": "Invalid credentials"})
}

func authStatusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "Auth service is running"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth/status", authStatusHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
