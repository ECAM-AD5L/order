package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"order/db"
	"order/schema"
	"strings"

	"github.com/gorilla/mux"
)

func createOrder(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	// check if user exist
	_, err := makeRequest(nil, fmt.Sprintf("http://user.ad5l.ecam.be/api/user/%s", name), "get")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var order schema.Order

	err = decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	/*
		for _, item := range order.Items {
			resp, err := makeRequest(nil, fmt.Sprintf("http://item.ad5l.ecam.be/api/item/get/%s", item.ID), "get")
			if err != nil {
				http.Error(w, "Unauthorized user", http.StatusUnauthorized)
			}
		}*/

	order.Status = "completed"

	err = db.CreateOrder(r.Context(), &order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	obj, err := db.GetOrder(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	obj, err := db.GetOrders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func UserHasItem(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]
	itemId := mux.Vars(r)["item_id"]

	obj, err := db.UserHasItem(r.Context(), userId, itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func makeRequest(message map[string]interface{}, url, method string) (map[string]interface{}, error) {

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	request, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
