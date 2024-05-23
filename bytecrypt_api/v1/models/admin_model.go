package models

import (
	"bytecrypt_api/utils"
	"fmt"

	"github.com/google/uuid"
)

type Administrator struct {
	Id       string
	Username string
	Name     string
	Email    string
	Password string
	Role     utils.Role
}

type AdminJson struct {
	Id       string
	Username string
	Name     string
	Email    string
	Password string
	Role     int32
}

func NewAdmin(username, name, email, password string, role utils.Role) Administrator {
	return Administrator{
		Id:       uuid.NewString(),
		Username: username,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func BlankAdmin() Administrator {
	return Administrator{}
}

func ValidateAdmin(admin *Administrator) error {
	var err error
	if admin.Id == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty id")
	}
	if admin.Email == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty email")
	}
	if admin.Name == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty name")
	}
	if admin.Username == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty username")
	}
	if admin.Password == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty password")
	}
	if admin.Role == utils.InvalidRole {
		err = fmt.Errorf("%w; %s", err, "cannot provide an erroneous role")
	}

	if err != nil {
		return fmt.Errorf("administrator validation errors: %s", err.Error()[2:])
	}
	return err
}

func (admin *Administrator) ToString() string {
	return fmt.Sprintf("ID: %s\n\tUsername: %s\n\tName: %s\n\tEmail: %s\n\tRole: %s",
		admin.Id, admin.Username, admin.Name, admin.Email, utils.RoleMap[admin.Role])
}
