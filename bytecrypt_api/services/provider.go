package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Provider struct {
	Queries *database.Queries
	Conn    *pgx.Conn
}

func NewProvider(backend *utils.Backend) (Provider, error) {
	queries, conn, err := database.NewDatabaseConnection(backend)
	if err != nil {
		return Provider{}, err
	}
	return Provider{Queries: queries, Conn: conn}, nil
}

func (provider *Provider) CloseDatabaseConnection(backend *utils.Backend) error {
	if provider.Conn != nil {
		err := provider.Conn.Close(context.Background())
		if err != nil {
			backend.Output <- fmt.Sprintf("Failed to close connection to database: %v", err)
			return err
		}
		backend.Output <- "Database connection successfully closed\n"
	}
	return nil
}
