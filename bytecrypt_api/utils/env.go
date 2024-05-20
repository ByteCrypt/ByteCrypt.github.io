package utils

type Environment string

const (
	JwtSecret        Environment = "JWT_SECRET"
	LogFile          Environment = "LOG_FILE"
	DbUri            Environment = "DB_URI"
	PostgresPassword Environment = "POSTGRES_PASSWORD"
)
