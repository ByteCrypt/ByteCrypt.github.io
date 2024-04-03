package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Subscription struct {
	Email string `json:"email"`
}

func subscribeHandler(writer http.ResponseWriter, req *http.Request) {
	var sub Subscription

	err := json.NewDecoder(req.Body).Decode(&sub)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Subscription request received for email: %s\n", sub.Email)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Subscription successful!"})
}

func main() {
	router := mux.NewRouter()

	// Set up CORS Middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if req.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(writer, req)
		})
	}

	// Apply CORS to all routes
	router.Use(corsMiddleware)
	router.HandleFunc("/api/subscribe", subscribeHandler).Methods("POST")

	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(writer, "Hellu, whale cum tu da weetahdit wunz")
	}).Methods("GET")

	fmt.Println("Server starting on port 5150...")
	if err := http.ListenAndServe(":5150", router); err != nil {
		fmt.Println("Error starting on server: ", err)
	}
}
