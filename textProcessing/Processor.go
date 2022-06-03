package textProcessing

import (
	"sort"
	"strings"
)

func Analyse(text string) (topWords []string, countWords int) {
	words := getParsedWords(text)

	keys := sortWordsByOccurrences(words)

	return keys, len(words)
}

func getParsedWords(text string) []string {
	return strings.Fields(text)
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
		normalizeWord := normalizeWord(word)

		occurrences, found := wordsMap[normalizeWord]
		if found {
			wordsMap[normalizeWord] = occurrences + 1
		} else {
			wordsMap[normalizeWord] = 1
			keys = append(keys, normalizeWord)
		}
	}
	return wordsMap, keys
}

func normalizeWord(word string) string {
	replacer := strings.NewReplacer(",", "", ".", "")

	wordNormalized := strings.ToLower(word)

	return replacer.Replace(wordNormalized)
}
