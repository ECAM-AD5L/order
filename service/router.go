package service

import (
	"github.com/gorilla/mux"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/order/{id}", getOrder).Methods("GET")
	return
}
