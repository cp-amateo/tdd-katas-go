package textProcessing

import (
	"reflect"
	"testing"
)

func Test_analyse(t *testing.T) {
	tests := []struct {
		name           string
		arg            string
		wantTopWords   []string
		wantCountWords int
	}{
		{"Analyse Hello", "Hello", []string{"hello"}, 1},
		{"Analyse Hello world", "Hello world", []string{"hello", "world"}, 2},
		{"Analyse Hello world world", "Hello world world", []string{"world", "hello"}, 3},
		{"Analyse full example", "Hello, world world world. Hello folk", []string{"world", "hello", "folk"}, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotTopWords, gotCountWords := Analyse(test.arg)
			if !reflect.DeepEqual(test.wantTopWords, gotTopWords) {
				t.Errorf("Analyse() gotTopWords = %v, want %v", gotTopWords, test.wantTopWords)
			}
			if test.wantCountWords != gotCountWords {
				t.Errorf("Analyse() gotCountWords = %v, want %v", gotCountWords, test.wantCountWords)
			}
		})
	}

}
