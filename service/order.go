package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order/db"
	"order/schema"

	"github.com/gorilla/mux"
)

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var item schema.Order

	err := decoder.Decode(&item)
	if err != nil {
		panic(err)
	}
	err = db.CreateOrder(r.Context(), &item)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	obj, err := db.GetOrder(r.Context(), id)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj, id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj, err := db.GetOrders(r.Context())
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func ListOrderByCustomerID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	obj, err := db.ListOrderByCustomerID(r.Context(), id)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj, id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}
