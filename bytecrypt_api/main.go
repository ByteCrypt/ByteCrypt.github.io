package main

import (
	"bytecrypt_api/services"
	"bytecrypt_api/utils"
	"fmt"
	"net/http"
	"strings"

	v1_controllers "bytecrypt_api/v1/controllers"
)

func main() {
	fmt.Println("Initializing backend")
	backend := utils.InitBackend()
	backend.Output <- "Backend initialized\n"
	defer backend.Shutdown()
	go processInput(backend)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/subscribe", v1_controllers.SubscribeHandler(backend))
	mux.HandleFunc("/api/v1/unsubscribe", v1_controllers.UnsubscribeHandler(backend))
	server := &http.Server{
		Addr:    ":5150",
		Handler: mux,
	}

	backend.Server = server

	go func() {
		if err := http.ListenAndServe(":5150", mux); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting on server: ", err)
			backend.Cancel()
		}
	}()

	<-backend.Context.Done()
}

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
			backend.Output <- "    - Seed: Seed operations"
			backend.Output <- "        [seed]"
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

	case "seed":
		if len(inputList) < 3 {
			backend.Output <- "Insufficient number of items passed:"
			backend.Output <- "  Expected: database seed [argument]"
			backend.Output <- "  - For help, use: database seed help"
			return
		}

		switch inputList[2] {
		case "help", "-h":
			{
				backend.Output <- "Use:"
				backend.Output <- "  database seed [argument]"
				backend.Output <- "  - Argument List:"
				backend.Output <- "    - Help: Shows this list"
				backend.Output <- "        [help | -h]"
				backend.Output <- "    - Initialize: Seed the database with initial values"
				backend.Output <- "        [initialize | init | -i | add | -a]"
				backend.Output <- "    - Delete: Remove the seeded values from the database"
				backend.Output <- "        [delete | -d | remove | -r]"
				backend.Output <- "    - List: List all seeded values in the database"
				backend.Output <- "        [list | ls | -l]"
			}

		case "initialize", "init", "-i", "add", "-a":
			{
				backend.Output <- "\n-- Test: Attempting to seed the database"
				provider, err := services.NewProvider(backend)
				if err != nil {
					backend.Output <- fmt.Sprintf("Could not create connection to database: %s", err.Error())
					return
				}
				defer provider.CloseDatabaseConnection(backend)

				err = provider.SeedDatabase()
				if err != nil {
					backend.Output <- err.Error()
					return
				}
				backend.Output <- "-- Test: Database seeded successfully"
			}

		case "delete", "-d", "remove", "-r":
			{
				backend.Output <- "\n-- Test: Attempting to remove database seed"
				provider, err := services.NewProvider(backend)
				if err != nil {
					backend.Output <- fmt.Sprintf("Could not create connection to database: %s", err.Error())
					return
				}
				defer provider.CloseDatabaseConnection(backend)

				err = provider.DeleteDatabaseSeed()
				if err != nil {
					backend.Output <- err.Error()
					return
				}
				backend.Output <- "-- Test: Database seed successfully removed"
			}

		case "list", "ls", "-l":
			{
				{
					backend.Output <- "\n-- Test: Attempting to remove database seed"
					provider, err := services.NewProvider(backend)
					if err != nil {
						backend.Output <- fmt.Sprintf("Could not create connection to database: %s", err.Error())
						return
					}
					defer provider.CloseDatabaseConnection(backend)

					seeds, err := provider.GetAllSeeded()
					if err != nil {
						backend.Output <- err.Error()
					}
					for i, seed := range seeds {
						backend.Output <- fmt.Sprintf("Seed [%d] == Name: %s -- Email: %s", i+1, seed.Name, seed.Email)
					}
				}
			}

		default:
			backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[2])
			backend.Output <- "  Expected: [help | add | delete | list]"
			backend.Output <- "  - For help, use: database seed help"
		}

	default:
		backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[1])
		backend.Output <- "  Expected [help | seed | test]"
		backend.Output <- "  - For help, use: database help"
	}
}
