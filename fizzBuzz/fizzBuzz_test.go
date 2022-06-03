package fizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("convert_1_to_1", func(t *testing.T) {
		got := FizzBuzz(1)
		want := "1"

		if want != got {
			t.Errorf("expectec FizzBuzz 1, got: %s", got)
		}
	})
}
