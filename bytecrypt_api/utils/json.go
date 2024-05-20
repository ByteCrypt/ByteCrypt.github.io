package utils

type JsonHeader string

const (
	ContentType     JsonHeader = "Content-Type"
	ApplicationJson JsonHeader = "application/json"
	Allow           JsonHeader = "Allow"
	Get             JsonHeader = "GET"
	Post            JsonHeader = "POST"
	Put             JsonHeader = "PUT"
)
