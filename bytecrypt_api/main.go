package main

import (
	"bytecrypt_api/common"
	"bytecrypt_api/database"
	"bytecrypt_api/services"
	"fmt"
	"net/http"
	"strings"

	v1_controllers "bytecrypt_api/v1/controllers"
)

func main() {
	fmt.Println("Initializing backend")
	backend := common.InitBackend()
	backend.Output <- "Backend initialized\n"
	defer backend.Shutdown()
	go processInput(backend)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/subscribe", v1_controllers.SubscribeHandler(backend))
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

func processInput(backend *common.Backend) {
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
			case "database", "db":
				processInputDatabase(backend, inputList)

			case "exit":
				backend.Cancel()

			default:
				backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[0])
				backend.Output <- "  Expected: [database | exit]"
			}
		}
	}
}

func processInputDatabase(backend *common.Backend, inputList []string) {
	if len(inputList) < 2 {
		backend.Output <- "Insufficient arguments passed"
		backend.Output <- "  Expected: database [argument]"
		return
	}

	switch inputList[1] {
	case "test":
		{
			backend.Output <- "\n-- Text: Testing database connection..."
			queries, conn, err := database.NewDatabaseConnection(backend)
			if err != nil {
				backend.Output <- "Could not create connection to database"
				return
			}
			provider := services.NewProvider(queries, conn)
			defer provider.CloseDatabaseConnection(backend)

			backend.Output <- "-- Test: Database connection successfully established"
		}

	case "seed":
		if len(inputList) < 3 {
			backend.Output <- "Insufficient number of items passed:"
			backend.Output <- "  Expected: database seed [argument]"
			return
		}

		switch inputList[2] {
		case "add", "init", "initialize":
			{
				backend.Output <- "\n-- Test: Attempting to seed the database"
				queries, conn, err := database.NewDatabaseConnection(backend)
				if err != nil {
					backend.Output <- "Could not create connection to database"
					return
				}
				provider := services.NewProvider(queries, conn)
				defer provider.CloseDatabaseConnection(backend)

				err = provider.SeedDatabase()
				if err != nil {
					backend.Output <- err.Error()
					return
				}
				backend.Output <- "-- Test: Database seeded successfully"
			}

		case "delete", "remove", "r", "d":
			{
				backend.Output <- "\n-- Test: Attempting to remove database seed"
				queries, conn, err := database.NewDatabaseConnection(backend)
				if err != nil {
					backend.Output <- "Could not create connection to database"
					return
				}
				provider := services.NewProvider(queries, conn)
				defer provider.CloseDatabaseConnection(backend)

				err = provider.DeleteDatabaseSeed()
				if err != nil {
					backend.Output <- err.Error()
					return
				}
				backend.Output <- "-- Test: Database seed successfully removed"
			}

		case "list", "ls":
			{
				{
					backend.Output <- "\n-- Test: Attempting to remove database seed"
					queries, conn, err := database.NewDatabaseConnection(backend)
					if err != nil {
						backend.Output <- "Could not create connection to database"
						return
					}
					provider := services.NewProvider(queries, conn)
					defer provider.CloseDatabaseConnection(backend)

					seeds, err := provider.GetAllSeeded()
					if err != nil {
						backend.Output <- err.Error()
					}
					for i, seed := range seeds {
						backend.Output <- fmt.Sprintf("Seed [%d] == Name: %s -- Email: %s", i, seed.Name, seed.Email)
					}
				}
			}

		default:
			backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[2])
			backend.Output <- "  Expected: [add | delete | list]"
		}

	default:
		backend.Output <- fmt.Sprintf("Invalid input: %s", inputList[1])
		backend.Output <- "  Expected [seed | test]"
	}
}
