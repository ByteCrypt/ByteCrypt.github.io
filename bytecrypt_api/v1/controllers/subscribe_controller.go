package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bytecrypt_api/common"
	"bytecrypt_api/database"
	"bytecrypt_api/services"
	"bytecrypt_api/v1/models"
)

func SubscribeHandler(backend *common.Backend) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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

				backend.Output <- fmt.Sprintf("Subscription request received for email: %s", sub.Email)

				queries, conn, err := database.NewDatabaseConnection(backend)
				if err != nil {
					backend.Output <- "Could not create connection to database\n"
					return
				}
				provider := services.NewProvider(queries, conn)
				defer provider.CloseDatabaseConnection(backend)

				provider.AddSubscription(sub)
				backend.Output <- fmt.Sprintf("%s now subscribed at %s", sub.Name, sub.Email)

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
}
