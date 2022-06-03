package fizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name       string
		arg        int
		wantResult string
	}{
		{"convert_1_to_1", 1, "1"},
		{"convert_2_to_2", 2, "2"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FizzBuzz(test.arg)
			if test.wantResult != got {
				t.Errorf("expectec FizzBuzz 1, got: %s", got)
			}
		})
	}
}
