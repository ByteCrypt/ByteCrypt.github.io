package common

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Backend struct {
	Output  chan string
	Input   chan string
	Context context.Context
	Cancel  context.CancelFunc
	Server  *http.Server
}

func InitBackend() *Backend {
	context, cancel := context.WithCancel(context.Background())
	backend := Backend{
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
	close(backend.Input)
	close(backend.Output)
}
