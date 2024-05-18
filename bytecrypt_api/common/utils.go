package common

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type BackEnd struct {
	Output  chan string
	Input   chan string
	Context context.Context
	Cancel  context.CancelFunc
}

func (backend *BackEnd) HandleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	go backend.ProcessInput()

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

func (backend *BackEnd) ProcessInput() {
	for text := range backend.Input {
		select {
		case <-backend.Context.Done():
			return

		default:
			{
				// Actually process the input text but for right now just print
				fmt.Println(text)
			}
		}
	}

}

func (backend *BackEnd) HandleOutput() {
	for output := range backend.Output {
		select {
		case <-backend.Context.Done():
			return

		default:
			{
				// Print the output
				fmt.Println(output)
			}
		}
	}
}
