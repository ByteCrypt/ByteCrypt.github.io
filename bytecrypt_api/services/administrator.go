package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/utils"
	"context"
)

type AddAdminParams struct {
	ID       string
	Email    string
	Name     string
	Username string
	Password string
	Role     int32
}

// TODO
func NewAddAdminParams(email, name, username, password string, role utils.Role) database.AddAdminParams {

	return database.AddAdminParams{}
}

func (provider *Provider) AddAdmin() error {

	provider.Queries.AddAdmin(context.Background(), database.AddAdminParams{})
	return nil
}
