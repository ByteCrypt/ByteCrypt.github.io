package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bytecrypt_api/v1/models"
)

func SubscribeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/api/v1/subscribe" {
		http.NotFound(writer, request)
		return
	}

	switch request.Method {
	case http.MethodPost:
		{
			var sub models.Subscription

			err := json.NewDecoder(request.Body).Decode(&sub)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}

			fmt.Printf("Subscription request received for email: %s\n", sub.Email)

			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(map[string]string{"message": "Subscription successful!"})
		}

	default:
		{
			writer.Header().Set("Allow", "POST")
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
