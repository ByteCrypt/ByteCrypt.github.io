package utils

type Role int32

const (
	InvalidRole Role = 0
	CoreAdmin   Role = 1
	Admin       Role = 2
	User        Role = 101
)

var RoleMap = map[Role]string{
	CoreAdmin: "Core Administrator",
	Admin:     "Administrator",
	User:      "User",
}

var Roles = []Role{
	CoreAdmin,
	Admin,
	User,
}
