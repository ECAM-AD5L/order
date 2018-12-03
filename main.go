package main

import (
	"fmt"
	"log"
	"net/http"
	"order/db"
	"order/service"
	"os"
	//"order/event"
)

func main() {
	// Create the database connection
	fmt.Println("Initializing the database...")
	repo, err := db.NewMongo()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db.SetRepository(repo)
	fmt.Println("Database initialized...")
	/*
		// Create the NATS connection
		es, err := event.NewNats()
		if err != nil {
			log.Println(err)
			panic(err)
		}
		event.SetEventStore(es)*/
	fmt.Println("Initializing the router...")
	router := service.NewRouter()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server running...")
}
