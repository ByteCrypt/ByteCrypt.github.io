package models

import (
	"bytecrypt_api/utils"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id       string
	Username string
	Name     string
	Email    string
	Password string
	Role     utils.Role
}

type UserJson struct {
	Username string
	Password string
}

func NewUser(username, name, email, password string, role utils.Role) User {
	return User{
		Id:       uuid.NewString(),
		Username: username,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}

func BlankUser() User {
	return User{}
}

func ValidateUser(user *User) error {
	var err error
	if user.Id == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty id")
	}
	if user.Email == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty email")
	}
	if user.Name == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty name")
	}
	if user.Username == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty username")
	}
	if user.Password == "" {
		err = fmt.Errorf("%w; %s", err, "cannot provide an empty password")
	}
	if user.Role == utils.InvalidRole {
		err = fmt.Errorf("%w; %s", err, "cannot provide an erroneous role")
	}

	if err != nil {
		return fmt.Errorf("administrator validation errors: %s", err.Error()[2:])
	}
	return err
}

func (user *User) ToString() string {
	return fmt.Sprintf("ID: %s\n\tUsername: %s\n\tName: %s\n\tEmail: %s\n\tRole: %s",
		user.Id, user.Username, user.Name, user.Email, utils.RoleMap[user.Role])
}
