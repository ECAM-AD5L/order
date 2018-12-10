package service

import (
	"encoding/json"
	"net/http"
	"order/db"
	"order/schema"

	"github.com/gorilla/mux"
)

func createOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var item schema.Order

	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = db.CreateOrder(r.Context(), &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	obj, err := db.GetOrder(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	obj, err := db.GetOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func ListOrderByCustomerID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	obj, err := db.ListOrderByCustomerID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}
