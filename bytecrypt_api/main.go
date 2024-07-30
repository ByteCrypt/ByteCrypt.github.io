package main

import (
	"bytecrypt_api/services"
	"bytecrypt_api/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	v1_controllers "bytecrypt_api/v1/controllers"
)

func main() {
	fmt.Println("Initializing backend")
	backend := utils.InitBackend()
	backend.Output <- "Backend initialized\n"
	defer backend.Shutdown()
	go processInput(backend)

	initServer(backend)

	<-backend.Context.Done()
}

func initServer(backend *utils.Backend) {
	// Init Mux
	mux := http.NewServeMux()

	// ------ Set up path functions
	mux.HandleFunc(string(v1_controllers.SubscribePath), v1_controllers.SubscribeHandler(backend))
	mux.HandleFunc(string(v1_controllers.UnsubscribePath), v1_controllers.UnsubscribeHandler(backend))
	mux.HandleFunc(string(v1_controllers.LoginPath), v1_controllers.LoginHandler(backend))

	server := &http.Server{
		Addr:    os.Getenv(string(utils.BackendAddress)),
		Handler: mux,
	}
	backend.Server = server
	go func() {
		if err := http.ListenAndServe(os.Getenv(string(utils.BackendAddress)), mux); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting on server: ", err)
			backend.Cancel()
		}
	}()
}

// Processes the input given by the user from the backend input buffer
func processInput(backend *utils.Backend) {
	for text := range backend.Input {
		select {
		case <-backend.Context.Done():
			return

		default:
			text = strings.ToLower(text)
			inputList := strings.Split(text, " ")
			if len(inputList) < 1 {
				continue
			}

			switch inputList[0] {
			case "help", "-h":
				{
					backend.Output <- "Use:"
					backend.Output <- "  [arguments]"
					backend.Output <- "  - Argument List:"
					backend.Output <- "    - Help:"
					backend.Output <- "        [help | -h]"
					backend.Output <- "    - Database: Database operations"
					backend.Output <- "        [database | db]"
					backend.Output <- "    - Exit: Shuts down the backend"
					backend.Output <- "        [exit]"

				}
			case "database", "db":
				processInputDatabase(backend, inputList, text)

			case "exit":
				backend.Cancel()

			default:
				backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[0])
				backend.Output <- "  Expected: [helps | database | exit]"
				backend.Output <- "  - For help, use: help"
			}
		}
	}
}

// Processes all database commands given by the user
func processInputDatabase(backend *utils.Backend, inputList []string, input string) {
	if len(inputList) < 2 {
		backend.Output <- "Insufficient arguments passed"
		backend.Output <- "  Expected: database [arguments]"
		backend.Output <- "  - For help, use: database help"
		return
	}

	switch inputList[1] {
	case "help", "-h":
		{
			backend.Output <- "Use:"
			backend.Output <- "  database [argument]"
			backend.Output <- "  - Argument List:"
			backend.Output <- "    - Help:"
			backend.Output <- "        [help | -h]"
			backend.Output <- "    - Test: Tests the connection to the database"
			backend.Output <- "        [test | -t]"
			backend.Output <- "    - Query (UNDER CONSTRUCTION): Perform a database query"
			backend.Output <- "        [query | -q]"
		}
	case "test", "-t":
		{
			backend.Output <- "\n-- Text: Testing database connection..."
			provider, err := services.NewProvider(backend)
			if err != nil {
				backend.Output <- fmt.Sprintf("Could not create connection to database: %s", err.Error())
				return
			}
			defer provider.CloseDatabaseConnection(backend)

			backend.Output <- "-- Test: Database connection successfully established"
		}

	case "query", "-q":
		{
			if len(inputList) < 3 {
				backend.Output <- "Insufficient arguments passed"
				backend.Output <- "  Expected: database query [arguments]"
			}

			switch inputList[2] {
			case "help", "-h":
				{
					backend.Output <- "Under construction"
				}

			default:
				{
					backend.Output <- "Under construction"
					backend.Output <- input
				}
			}
		}

	default:
		backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[1])
		backend.Output <- "  Expected [help | seed | test]"
		backend.Output <- "  - For help, use: database help"
	}
}
