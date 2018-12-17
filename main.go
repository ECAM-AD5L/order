package main

import (
	"fmt"
	"log"
	"order/db"
	"order/service"
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
	fmt.Println("Initializing the server...")
	service.StartAPIServer()
	fmt.Println("Server running...")
}
