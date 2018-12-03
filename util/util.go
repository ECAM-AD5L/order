package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func ResponseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}

func GetDBConnection() (*mongo.Client, error) {
	return mongo.NewClient(connectionString())
}

func connectionString() string {
	host := os.Getenv("MONGO_URL")
	return fmt.Sprintf("%s", host)
}
