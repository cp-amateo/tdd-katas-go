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
		{"analyse Hello", "Hello", []string{"Hello"}, 1},
		{"analyse Hello world", "Hello", []string{"Hello", "world"}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotTopWords, gotCountWords := analyse(test.arg)
			if !reflect.DeepEqual(test.wantTopWords, gotTopWords) {
				t.Errorf("analyse() gotTopWords = %v, want %v", gotTopWords, test.wantTopWords)
			}
			if test.wantCountWords != gotCountWords {
				t.Errorf("analyse() gotCountWords = %v, want %v", gotCountWords, test.wantCountWords)
			}
		})
	}

}
