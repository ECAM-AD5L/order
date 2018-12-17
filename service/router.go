package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	router.HandleFunc("/orders/me/{id}", ListOrderByCustomerID).Methods("GET")
	router.HandleFunc("/orders/{user_id}/{item_id}", UserHasItem).Methods("GET")

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
