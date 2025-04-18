package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

var orders []Order

func createOrder(w http.ResponseWriter, r *http.Request) {
	var o Order
	json.NewDecoder(r.Body).Decode(&o)
	o.ID = len(orders) + 1
	orders = append(orders, o)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

func listOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func orderStatusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "Order service is running"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", createOrder).Methods("POST")
	r.HandleFunc("/orders", listOrders).Methods("GET")
	r.HandleFunc("/orders/status", orderStatusHandler).Methods("GET")
	http.ListenAndServe(":8082", r)
}
