package main

import (
	"fmt"
	"net/http"

	v1_controllers "bytecrypt_api/v1/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/api/v1/subscribe", v1_controllers.SubscribeHandler)
	if err := http.ListenAndServe(":5150", mux); err != nil {
		fmt.Println("Error starting on server: ", err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	switch request.Method {
	case http.MethodGet:
		{

		}

	case http.MethodPost:
		{

		}

	case http.MethodOptions:
		{

		}

	default:
		{
			writer.Header().Set("Allow", "GET, POST, OPTIONS")
			http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
