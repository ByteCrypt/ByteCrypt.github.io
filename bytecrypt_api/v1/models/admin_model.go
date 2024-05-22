package models

import (
	"bytecrypt_api/utils"

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
