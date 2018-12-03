package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"strconv"
	"strings"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", getOrder).Methods("GET")

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		fmt.Println("Error parsing the DEBUG env variable")
	}
	if debug {
		err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pathTemplate, err := route.GetPathTemplate()
			if err == nil {
				fmt.Print("ROUTE:", pathTemplate)
			}
			methods, err := route.GetMethods()
			if err == nil {
				fmt.Println(" Methods:", strings.Join(methods, ","))
			}
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
