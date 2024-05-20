package utils

type Level uint8

const (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3

	PERMISSIONS = 0666
)

type Log struct {
	Level   Level
	Message string
}

func NewLog(level Level, message string) Log {
	return Log{Level: level, Message: message}
}
