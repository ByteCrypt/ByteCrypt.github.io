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
		backend.Output <- fmt.Sprintf("Failed to connect to database: %v\nAttempting again in 10 seconds", err)
		time.Sleep(time.Second * 10)
	}

	queries := New(conn)

	err = runMigrations("./database/schemas/schema.sql", queries, &context)
	if err != nil {
		return nil, nil, err
	}

	return queries, conn, nil
}
