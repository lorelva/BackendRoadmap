package str

import (
	"testing"
)

func TestStr(t *testing.T) {
	t.Parallel()
	// "#"
	// 0, error
	t.Run("String Symbol to int error", func(t *testing.T) {
		result, err := ConvertString("#")
		if err == nil && result != 0 {
			t.Errorf("String to convert('#'), expected: 0, got: %v", result)
		}
	})
	t.Run("String word to int  error", func(t *testing.T) {
		result, err := ConvertString("word")
		if err == nil && result != 0 {
			t.Errorf("String to convert('word'), expected: 0, got: %v", result)
		}
	})
	t.Run("Convert number string to int ok", func(t *testing.T) {
		result, err := ConvertString("123")
		if err != nil && result != 123 {
			t.Errorf("String to convert('123'), expected: 123, got: %v", result)
		}
	})
}
