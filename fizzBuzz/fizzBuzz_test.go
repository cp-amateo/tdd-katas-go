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
		{"convert_3_to_Fizz", 3, "Fizz"},
		{"convert_5_to_Buzz", 5, "Buzz"},
		{"convert_15_to_FizzBuzz", 15, "FizzBuzz"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FizzBuzz(test.arg)
			if test.wantResult != got {
				t.Errorf("expectec FizzBuzz %s, got: %s", test.wantResult, got)
			}
		})
	}
}
