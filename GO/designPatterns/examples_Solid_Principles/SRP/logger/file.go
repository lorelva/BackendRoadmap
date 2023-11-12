package logger

import (
	"errors"
	"fmt"
	"log"
)

// logging to a file
type FileLogger struct {
}

// Log(message string)
func (f *FileLogger) Log(message string) {

	for i := 0; i < 10; i++ {
		// Log to a file
		log.Printf("El numero es: %d - %s \n ", i, message)

	}

	for i := 0; i < 20; i++ {
		if i < 4 {
			fmt.Println("adasdas")
		}
	}

	fmt.Println("dfasdsadas")

}

func (f *FileLogger) Clean(msg string) (bool, error) {
	//Mandar a función privada con metodo
	if msg != "" {

		return false, errors.New("mensaje no vacía")
	}

	return true, nil

}
