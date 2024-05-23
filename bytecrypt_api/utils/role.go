package utils

type Role int32

const (
	InvalidRole Role = 0
	CoreAdmin   Role = 1
)

var Roles = []Role{
	CoreAdmin,
}

var RoleMap = map[Role]string{
	CoreAdmin: "Core Administrator",
}
