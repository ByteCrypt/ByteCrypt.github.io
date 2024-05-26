package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/utils"
	"context"
	"fmt"
)

func InitRoles(backend *utils.Backend) error {
	provider, err := NewProvider(backend)
	if err != nil {
		return err
	}
	defer provider.CloseDatabaseConnection(backend)

	var errs error
	for _, role := range utils.Roles {
		if _, err := provider.Queries.GetRoleById(context.Background(), int32(role)); err == nil {
			continue
		}
		add := database.AddRoleParams{
			ID:    int32(role),
			Title: utils.RoleMap[role],
		}
		_, err := provider.Queries.AddRole(context.Background(), add)
		if err != nil {
			errs = fmt.Errorf("%w; %s", errs, err)
		}
	}
	if errs != nil {
		return fmt.Errorf("%s", errs.Error()[:2])
	}

	return nil
}

func NewRoleParams(role utils.Role) (database.AddRoleParams, error) {
	title := utils.RoleMap[role]
	if len(title) == 0 {
		return database.AddRoleParams{}, fmt.Errorf("role not found in map")
	}

	return database.AddRoleParams{
		ID:    int32(role),
		Title: title,
	}, nil
}

func (provider *Provider) GetRoleById(id int32) (utils.Role, error) {
	role, err := provider.Queries.GetRoleById(context.Background(), id)
	if err != nil {
		return utils.InvalidRole, err
	}

	return utils.Role(role.ID), nil
}

func (provider *Provider) GetRoleByTitle(title string) (utils.Role, error) {
	role, err := provider.Queries.GetRoleByTitle(context.Background(), title)
	if err != nil {
		return utils.InvalidRole, err
	}

	return utils.Role(role.ID), nil
}

func (provider *Provider) AddRole(role utils.Role) error {

	return nil
}
