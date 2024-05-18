package main

import (
	"bytecrypt_api/common"
	"context"
	"fmt"
	"net/http"

	v1_controllers "bytecrypt_api/v1/controllers"
)

func main() {
	context, cancel := context.WithCancel(context.Background())
	backend := common.BackEnd{
		Output:  make(chan string, 100),
		Input:   make(chan string, 100),
		Context: context,
		Cancel:  cancel,
	}
	go backend.HandleOutput()
	go backend.HandleInput()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/subscribe", v1_controllers.SubscribeHandler)
	if err := http.ListenAndServe(":5150", mux); err != nil {
		fmt.Println("Error starting on server: ", err)
	}
}
