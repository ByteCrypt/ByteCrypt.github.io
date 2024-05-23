package controllers

type Path string

const (
	SubscribePath   Path = "/api/v1/subscribe"
	UnsubscribePath Path = "/api/v1/unsubscribe"
	AdminLoginPath  Path = "api/v1/admin_login"
)
