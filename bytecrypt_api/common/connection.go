package common

import (
	"bytecrypt_api/database"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type Provider struct {
	Queries *database.Queries
}

func NewProvider(queries *database.Queries) Provider {
	return Provider{Queries: queries}
}

func (backend *BackEnd) NewDatabaseConnection() (*database.Queries, error) {
	context := context.Background()
	connString := os.Getenv("DB_URI")
	if len(connString) == 0 {
		return nil, errors.New("no database uri found in environment variables")
	}

	var (
		conn *pgx.Conn
		err  error
	)
	// Attempt 5 times
	for i := 1; i <= 5; i++ {
		backend.Output <- fmt.Sprintf("[TRY %d] Attempting to connecting to %v", i, connString)
		conn, err = pgx.Connect(context, connString)
		if err == nil {
			backend.Output <- "Successfully connected to database"
			break
		}
		backend.Output <- fmt.Sprintf("Failed to connect to database: %v\nAttempting again in 10 seconds", err)
		time.Sleep(time.Second * 10)
	}

	queries := database.New(conn)

	// Migrations here

	return queries, nil
}
