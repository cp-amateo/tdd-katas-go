package textProcessing

import "strings"

func analyse(text string) (topWords []string, countWords int) {
	words := strings.Fields(text)
	return words, len(words)
}
