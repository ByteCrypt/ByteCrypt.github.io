package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/utils"
	"bytecrypt_api/v1/models"
	"context"
	"errors"
)

func NewAddAdminParams(id, email, name, username, password string, role utils.Role) database.AddAdminParams {
	return database.AddAdminParams{
		ID:       id,
		Email:    email,
		Name:     name,
		Username: username,
		Password: password,
		Role:     int32(role),
	}
}

func ConvertAdminToDb(admin models.Administrator) database.AddAdminParams {
	return database.AddAdminParams{
		ID:       admin.Id,
		Email:    admin.Email,
		Name:     admin.Name,
		Username: admin.Username,
		Password: admin.Password,
		Role:     int32(admin.Role),
	}
}

func ConvertDbToAdmin(admin database.Administrator) models.Administrator {
	return models.Administrator{
		Id:       admin.ID,
		Email:    admin.Email,
		Name:     admin.Name,
		Username: admin.Username,
		Password: admin.Password,
		Role:     utils.Role(admin.Role),
	}
}

func (provider *Provider) AddAdmin(admin models.Administrator) (models.Administrator, error) {
	if err := provider.ValidateEmail(admin.Email); err != nil {
		return admin, err
	}

	if dbadmin, err := provider.GetAdminByEmail(admin.Email); err == nil {
		errors.Join(err, errors.New("an administrator with that email already exists"))
		if dbadmin.Username == admin.Username {
			errors.Join(err, errors.New("an administrator with that username already exists"))
		}
		if dbadmin.Id == admin.Id {
			errors.Join(err, errors.New("an administrator with that id already exists"))
		}
		return admin, err
	}

	if _, err := provider.Queries.AddAdmin(context.Background(), ConvertAdminToDb(admin)); err != nil {
		return admin, err
	}

	return admin, nil
}

func (provider *Provider) GetAdminById(id string) (models.Administrator, error) {
	admin, err := provider.Queries.GetAdminById(context.Background(), id)
	if err != nil {
		return models.BlankAdmin(), err
	}

	return ConvertDbToAdmin(admin), nil
}

func (provider *Provider) GetAdminByEmail(email string) (models.Administrator, error) {
	if err := provider.ValidateEmail(email); err != nil {
		return models.BlankAdmin(), err
	}

	admin, err := provider.Queries.GetAdminByEmail(context.Background(), email)
	if err != nil {
		return models.BlankAdmin(), err
	}

	return ConvertDbToAdmin(admin), nil
}

func (provider *Provider) GetAdminByUsername(username string) (models.Administrator, error) {
	admin, err := provider.Queries.GetAdminByUsername(context.Background(), username)
	if err != nil {
		return models.BlankAdmin(), err
	}

	return ConvertDbToAdmin(admin), nil
}
