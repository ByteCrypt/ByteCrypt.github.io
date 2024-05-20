package utils

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Backend struct {
	File    *os.File
	Logger  *slog.Logger
	Log     chan Log
	Output  chan string
	Input   chan string
	Context context.Context
	Cancel  context.CancelFunc
	Server  *http.Server
}

func InitBackend() *Backend {
	file, err := os.OpenFile(os.Getenv(string(LogFile)), os.O_CREATE|os.O_WRONLY|os.O_APPEND, PERMISSIONS)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(file, nil))
	slog.SetDefault(logger)

	context, cancel := context.WithCancel(context.Background())
	backend := Backend{
		File:    file,
		Logger:  logger,
		Log:     make(chan Log, 100),
		Output:  make(chan string, 100),
		Input:   make(chan string, 100),
		Context: context,
		Cancel:  cancel,
	}
	backend.init()
	return &backend
}

func (backend *Backend) init() {
	go backend.HandleInput()
	go backend.HandleOutput()
	go backend.HandleLog()
}

func (backend *Backend) HandleInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		select {
		case <-backend.Context.Done():
			return

		default:
			{
				backend.Input <- scanner.Text()
			}
		}
	}
}

func (backend *Backend) HandleOutput() {
	for output := range backend.Output {
		select {
		case <-backend.Context.Done():
			return

		default:
			{
				fmt.Println(output)
			}
		}
	}
}

func (backend *Backend) HandleLog() {
	for log := range backend.Log {
		select {
		case <-backend.Context.Done():
			return

		default:
			{
				switch log.Level {
				case Debug:
					slog.Debug(log.Message)

				case Info:
					slog.Info(log.Message)

				case Warn:
					slog.Warn(log.Message)

				case Error:
					slog.Error(log.Message)

				default:
					backend.Output <- "Invalid log level passed"
				}
			}
		}
	}
}

func (backend *Backend) Shutdown() {
	fmt.Println("Exiting backend")

	backend.Cancel()
	if backend.Server != nil {
		context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := backend.Server.Shutdown(context); err != nil {
			fmt.Println("Error shutting down server:", err)
		}
	}
	// Close all buffers
	close(backend.Input)
	close(backend.Output)
	close(backend.Log)

	// Empty the log buffer
	for log := range backend.Log {
		switch log.Level {
		case Debug:
			slog.Debug(log.Message)

		case Info:
			slog.Info(log.Message)

		case Warn:
			slog.Warn(log.Message)

		case Error:
			slog.Error(log.Message)
		}
	}

	// Close the file
	backend.File.Close()
}
