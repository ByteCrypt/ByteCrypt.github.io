package controllers

type Path string

const (
	SubscribePath   Path = "/api/v1/subscribe"
	UnsubscribePath Path = "/api/v1/unsubscribe"
	LoginPath       Path = "/api/v1/login"
	RegisterPath    Path = "/api/v1/register"
)
