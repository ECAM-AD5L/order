package main

import (
	"fmt"
	"log"
	"order/db"
	"order/service"

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
	fmt.Println("Initializing the server...")
	service.StartAPIServer()
	fmt.Println("Server running...")
}
