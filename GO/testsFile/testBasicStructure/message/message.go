package message

import (
	"errors"
	"fmt"
)

// Karina
// -> Hello Karina
func Message(name string) (string, error) {
	if name == "" {
		return "", errors.New("cannot send empty name")
	}

	return fmt.Sprintf("Hello %v", name), nil
}
