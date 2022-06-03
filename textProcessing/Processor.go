package textProcessing

import (
	"sort"
	"strings"
)

func analyse(text string) (topWords []string, countWords int) {
	words := strings.Fields(text)

	wordsMap := make(map[string]int)

	for _, word := range words {
		occurrences, found := wordsMap[word]
		if found {
			wordsMap[word] = occurrences + 1
		} else {
			wordsMap[word] = 1
		}
	}

	keys := make([]string, 0, len(wordsMap))

	for key := range wordsMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return wordsMap[keys[i]] > wordsMap[keys[j]]
	})

	return keys, len(words)
}
