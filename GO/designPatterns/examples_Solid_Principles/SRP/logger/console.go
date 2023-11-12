package logger

import (
	"errors"
	"log"
)

// ConsoleLogger and FileLogger have single responsibilities:
// logging to the console
type ConsoleLogger struct {
}

func (c *ConsoleLogger) Log(message string) {
	// Log to the console
	log.Printf("El mensaje de consola es: %s\n", message)
}

func (c *ConsoleLogger) Clean(msg string) (bool, error) {
	if msg == "" {

		return false, errors.New("mensaje vacío")
	}

	return true, nil

}
