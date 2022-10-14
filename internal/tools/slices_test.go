package tools

import "testing"

func TestContains(t *testing.T) {
	t.Run("An item not in the array generates a False", func(t *testing.T) {
		value := "first"
		slice := []any{"second", "third", "fourth", "fifth"}

		result := Contains(value, slice)

		if result {
			t.Error("We expected a false value here")
		}
	})

	t.Run("An item in the array generates a True", func(t *testing.T) {
		value := "first"
		slice := []any{"first", "second", "third", "fourth", "fifth"}

		result := Contains(value, slice)

		if !result {
			t.Error("We expected a true value here")
		}
	})
}
