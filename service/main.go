package main

import (
	"log"
	"net/http"
	"order/db"
	"order/event"

	"github.com/gorilla/mux"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/order/{id}", getOrder).Methods("GET")
	return
}

func main() {
	// Create the database connection
	repo, err := db.NewMongo()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db.SetRepository(repo)

	// Create the NATS connection
	es, err := event.NewNats()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	event.SetEventStore(es)

	router := newRouter()

	if err := http.ListenAndServe(":9090", router); err != nil {
		log.Fatal(err)
	}
}
