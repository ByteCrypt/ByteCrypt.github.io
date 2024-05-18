package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bytecrypt_api/common"
	"bytecrypt_api/database"
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

// ############################### Seed ###############################
func SeedDatabase(queries *database.Queries) error {
	seedList := []models.Subscription{
		{Email: "a@example.mail", Name: "a"},
		{Email: "b@example.mail", Name: "b"},
		{Email: "c@example.mail", Name: "c"},
		{Email: "d@example.mail", Name: "t"},
		{Email: "e@example.mail", Name: "e"},
		{Email: "f@example.mail", Name: "f"},
		{Email: "g@example.mail", Name: "g"},
		{Email: "h@example.mail", Name: "h"},
		{Email: "i@example.mail", Name: "i"},
		{Email: "j@example.mail", Name: "j"},
		{Email: "k@example.mail", Name: "k"},
		{Email: "l@example.mail", Name: "l"},
		{Email: "m@example.mail", Name: "m"},
		{Email: "n@example.mail", Name: "n"},
		{Email: "o@example.mail", Name: "o"},
		{Email: "p@example.mail", Name: "p"},
		{Email: "q@example.mail", Name: "q"},
		{Email: "r@example.mail", Name: "r"},
		{Email: "s@example.mail", Name: "s"},
		{Email: "t@example.mail", Name: "t"},
		{Email: "u@example.mail", Name: "u"},
		{Email: "v@example.mail", Name: "v"},
		{Email: "w@example.mail", Name: "w"},
		{Email: "x@example.mail", Name: "x"},
		{Email: "y@example.mail", Name: "y"},
		{Email: "z@example.mail", Name: "z"},
	}

	provider := common.NewProvider(queries)

	for _, sub := range seedList {

	}

	return nil
}
