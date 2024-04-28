package suma

import "testing"

func TestSumar(t *testing.T) {
	// ejecuta en parelelo en la mayoria de casos
	// execute paralel almost in all cases
	t.Parallel()

	//The Run methods of T and B allow defining subtests
	t.Run("Add Digits ok", func(t *testing.T) {
		result := Add(1, 2)
		if result != 3 {
			t.Errorf("Add(1,2), expected: 3, got: %v", result)
		}
	})

	t.Run("Add negative digits ok", func(t *testing.T) {
		result := Add(10, -5)
		if result != 5 {
			t.Errorf("Add(10,-5), expected: 5, got: %v", result)
		}
	})

}

func TestDividir(t *testing.T) {
	t.Parallel()

	t.Run("Divide 0 error", func(t *testing.T) {
		result, err := Divide(5, 0)
		if err == nil && result != 0 {
			t.Errorf("Divide(5,0), expected: error, got %v. Result: %v", err, result)
		}
	})

	t.Run("Divide numbers OK", func(t *testing.T) {
		result, err := Divide(10, 2)
		if result != 5 && err != nil {
			t.Errorf("Divide(10,2), expected: 5, got %v, Error %v", result, err)
		}
	})

}
