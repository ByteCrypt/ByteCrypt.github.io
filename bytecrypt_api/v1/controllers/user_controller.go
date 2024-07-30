package controllers

import (
	"bytecrypt_api/services"
	"bytecrypt_api/utils"
	"bytecrypt_api/v1/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(backend *utils.Backend) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != string(LoginPath) {
			http.NotFound(writer, request)
			return
		}

		switch request.Method {
		case http.MethodPost:
			{
				var user models.UserJson

				err := json.NewDecoder(request.Body).Decode(&user)
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

				// Attempt the log in
				// Get the user
				level := utils.Info
				ok := true
				message := "Login attempt was successful"
				dbUser, err := provider.GetUserByUsername(user.Username)
				if err != nil {
					level = utils.Error
					message = "Could not find by that username"
					ok = false
				}

				backend.Log <- utils.NewLog(level, message)
				backend.Output <- dbUser.Id

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
