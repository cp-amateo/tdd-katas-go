package textProcessing

import (
	"sort"
	"strings"
)

func Analyse(text string) (topWords []string, countWords int) {
	words := strings.Fields(text)

	keys := sortWordsByOccurrences(words)

	return keys, len(words)
}

func sortWordsByOccurrences(words []string) []string {
	wordsMap, keys := createWordsMap(words)

	sort.SliceStable(keys, func(i, j int) bool {
		return wordsMap[keys[i]] > wordsMap[keys[j]]
	})
	return keys
}

func createWordsMap(words []string) (map[string]int, []string) {
	wordsMap := make(map[string]int)
	keys := make([]string, 0, len(words))

	for _, word := range words {
		word := strings.ToLower(word)
		replacer := strings.NewReplacer(",", "", ".", "")
		word = replacer.Replace(word)

		occurrences, found := wordsMap[word]
		if found {
			wordsMap[word] = occurrences + 1
		} else {
			wordsMap[word] = 1
			keys = append(keys, word)
		}
	}
	return wordsMap, keys
}
