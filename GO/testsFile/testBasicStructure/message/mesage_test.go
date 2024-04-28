package message

import "testing"

func TestMessage(t *testing.T) {
	t.Parallel()

	t.Run("Empty name for message error", func(t *testing.T) {
		result, err := Message("")
		expectedMessage := "cannot send empty name"
		if err == nil && result != expectedMessage {
			t.Errorf("Name for expected message: %s , got: %v", expectedMessage, result)
		}
	})
	t.Run("Name for message ok", func(t *testing.T) {
		result, err := Message("Maria")
		expectedMessage := "Hello Lorena"
		if err != nil && result != expectedMessage {
			t.Errorf("Name for expected message: %s, got: %v", expectedMessage, result)
		}
	})
}
