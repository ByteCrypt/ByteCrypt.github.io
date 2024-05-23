package controllers

import (
	"bytecrypt_api/services"
	"bytecrypt_api/utils"
	"bytecrypt_api/v1/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func AdminLoginHandler(backend *utils.Backend) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(AdminLoginPath) {
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

				provider, err := services.NewProvider(backend)
				if err != nil {
					backend.Log <- utils.NewLog(utils.Error, fmt.Sprintf("Could not create connection to database: %s", err.Error()))
					return
				}
				defer provider.CloseDatabaseConnection(backend)

				level := utils.Info
				ok := true
				message := "Subscription attempt was successful"
				_, err = provider.AddSubscription(sub)
				if err != nil {
					level = utils.Warn
					ok = false
					message = fmt.Sprintf("Subscription attempt was unsuccessful: %s", err.Error())
				}
				backend.Log <- utils.NewLog(level, fmt.Sprintf("Subscription Attempt: %s::%s -- %s", sub.Name, sub.Email, message))

				writer.Header().Set(string(utils.ContentType), string(utils.ApplicationJson))
				subJson := models.NewSubscriptionResponse(ok, message)
				json.NewEncoder(writer).Encode(subJson)
			}

		default:
			{
				writer.Header().Set(string(utils.Allow), string(utils.Post))
				http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
			}
		}
	}
}
