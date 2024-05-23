package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bytecrypt_api/services"
	"bytecrypt_api/utils"
	"bytecrypt_api/v1/models"
)

func SubscribeHandler(backend *utils.Backend) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(SubscribePath) {
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

func UnsubscribeHandler(backend *utils.Backend) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(UnsubscribePath) {
			http.NotFound(writer, request)
			return
		}

		switch request.Method {
		case http.MethodPost:
			{
				var sub models.Subscription
				if err := json.NewDecoder(request.Body).Decode(&sub); err != nil {
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
				message := "Successfully unsubscribed"
				ok := true
				if err := provider.RemoveSubscription(sub.Email); err != nil {
					level = utils.Warn
					ok = false
					message = fmt.Sprintf("Unable to unsuscribe: %s", err.Error())
				}
				backend.Log <- utils.NewLog(level, fmt.Sprintf("Unsubscription Attempt: %s::%s -- %s", sub.Name, sub.Email, message))

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
