package database

import (
	"bytecrypt_api/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

// Performs migrations of the database into the necessary migration file
func runMigrations(migrationFile string, queries *Queries, context *context.Context) error {
	fileContents, err := os.ReadFile(migrationFile)
	if err != nil {
		return err
	}

	_, err = queries.db.Exec(*context, string(fileContents))
	if err != nil {
		return err
	}

	return nil
}

func NewDatabaseConnection(backend *utils.Backend) (*Queries, *pgx.Conn, error) {
	context := context.Background()
	connString := os.Getenv(string(utils.DbUri))
	if len(connString) == 0 {
		return nil, nil, errors.New("no database uri found in environment variables")
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
		if i == 5 {
			backend.Log <- utils.NewLog(utils.Error, fmt.Sprintf("Could not connect to database: %v", err))
			return nil, nil, err
		}
		backend.Output <- fmt.Sprintf("Failed to connect to database: %v\nAttempting again in 10 seconds", err)
		backend.Log <- utils.NewLog(utils.Warn, fmt.Sprintf("Failed to connect to database, attempt %d", i))
		time.Sleep(time.Second * 10)
	}

	queries := New(conn)

	err = runMigrations(string(utils.Schema), queries, &context)
	if err != nil {
		backend.Log <- utils.NewLog(utils.Error, fmt.Sprintf("Failed to run migrations: %v", err))
		return nil, nil, err
	}

	return queries, conn, nil
}
