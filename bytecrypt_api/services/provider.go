package services

import (
	"bytecrypt_api/common"
	"bytecrypt_api/database"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Provider struct {
	Queries *database.Queries
	Conn    *pgx.Conn
}

func NewProvider(queries *database.Queries, conn *pgx.Conn) Provider {
	return Provider{Queries: queries, Conn: conn}
}

func (provider *Provider) CloseDatabaseConnection(backend *common.Backend) error {
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
